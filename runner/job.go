package runner

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)

func (r *Runner) CreateJob(job *pb.Job) (err error) {
	glog.Infof("runner %s creating job %d", r.ID, job.Id)
	if job.Type == pb.JobType_UNKNOWN_TYPE {
		err = errors.New(fmt.Sprintf("Runner CreateJob: invalid JobType: %v", job.Type))
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}
	// if eval job or infer job: down load model
	if job.Type == pb.JobType_INFER || job.Type == pb.JobType_EVALUATION {
		_, err = r.downloadAndUnarchiveModel(job)
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
			return
		}
	}
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}
	defer cli.Close()
	// pull image
	reader, err := cli.ImagePull(ctx, DEFAULT_IMAGE, types.ImagePullOptions{})
	if err != nil {
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}
	io.Copy(os.Stdout, reader)

	plugin := plugins.New(job)
	err = r.createShellFile(job, plugin)
	if err != nil {
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}
	jobIdStr := strconv.FormatUint(job.Id, 10)
	logFilePath := r.createLogFile(jobIdStr)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        DEFAULT_IMAGE,
		Cmd:          r.getCmd(path.Join(r.Conf.JobShellDir, jobIdStr)),
		Env:          []string{fmt.Sprintf("TUSHARE_TOKEN=%s", r.Conf.TushareToken)},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
	}, &container.HostConfig{
		Mounts: r.getTbaseMounts(jobIdStr),
	}, nil, jobIdStr)
	if err != nil {
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return
	}

	r.Containers[job.Id] = Container{
		Name:    strconv.FormatUint(job.Id, 10),
		ID:      resp.ID,
		ShortID: resp.ID[:12],
		Job:     job,
		Plugin:  plugin,
	}

	glog.Infof("runner %s created job %d, container id: %s", r.ID, job.Id, resp.ID)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		glog.Error(err)
		job.Status = pb.JobStatus_FAILED
		return err
	}

	// stream read container log and write into file
	go writeLog(resp.ID, logFilePath)

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
			err = r.StopJob(job)
			if err != nil {
				glog.Error(err)
				job.Status = pb.JobStatus_FAILED
				return err
			}
		}
	case <-statusCh:
		glog.Infof("runner %s completed job %d, container id: %s", r.ID, job.Id, resp.ID)
		err = r.StopJob(job)
		if err != nil {
			glog.Error(err)
			job.Status = pb.JobStatus_FAILED
			return err
		}
		job.Status = pb.JobStatus_SUCCESSED
		glog.Infof("runner %s clean job %d, container id: %s", r.ID, job.Id, resp.ID)
		return
	}
	return
}

func (r *Runner) StopJob(job *pb.Job) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		return
	}
	containerId := r.Containers[job.Id].ShortID
	glog.Infof("runner %s stopping job %d, container id: %s", r.ID, job.Id, containerId)
	if err := cli.ContainerStop(ctx, containerId, nil); err != nil {
		glog.Error(err)
		return err
	}
	glog.Infof("runner %s stopped job %d, container id: %s", r.ID, job.Id, containerId)
	err = r.RemoveContainer(job)
	if err != nil {
		glog.Error(err)
		return err
	}
	err = r.FinishedJob(job)
	if err != nil {
		glog.Error(err)
		return err
	}
	return
}

func (r *Runner) RemoveContainer(job *pb.Job) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		return
	}
	// remove container
	containerId := r.Containers[job.Id].ShortID
	glog.Infof("runner %s removing container id: %s, job id %d", r.ID, containerId, job.Id)
	if err := cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{}); err != nil {
		glog.Error(err)
		return err
	}
	glog.Infof("runner %s removed container id: %s, job %d", r.ID, containerId, job.Id)
	return
}
