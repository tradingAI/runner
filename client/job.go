package client

import (
	"context"
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

func (c *Client) CreateJob(job *pb.Job) (err error) {
	glog.Infof("runner %s creating job %d", c.ID, job.Id)
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		return
	}
	// pull image
	reader, err := cli.ImagePull(ctx, DEFAULT_IMAGE, types.ImagePullOptions{})
	if err != nil {
		glog.Error(err)
		return
	}
	io.Copy(os.Stdout, reader)

	plugin := plugins.New(job)
	err = c.createShellFile(job, plugin)
	if err != nil {
		glog.Error(err)
		return
	}
	jobIdStr := strconv.FormatUint(job.Id, 10)
	logFilePath := c.createLogFile(jobIdStr)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        DEFAULT_IMAGE,
		Cmd:          c.getCmd(path.Join(c.Conf.JobShellDir, jobIdStr)),
		Env:          []string{fmt.Sprintf("TUSHARE_TOKEN=%s", c.Conf.TushareToken)},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
	}, &container.HostConfig{
		Mounts: c.getTbaseMounts(jobIdStr),
	}, nil, jobIdStr)
	if err != nil {
		glog.Error(err)
		return
	}

	c.Containers[job.Id] = Container{
		Name:    strconv.FormatUint(job.Id, 10),
		ID:      resp.ID,
		ShortID: resp.ID[:12],
		Job:     job,
		Plugin:  plugin,
	}

	defer c.RemoveContainer(job.Id)

	glog.Infof("runner %s created job %d, container id: %s", c.ID, job.Id, resp.ID)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		glog.Error(err)
		return err
	}

	// stream read container log and write into file
	go writeLog(resp.ID, logFilePath)

	ch := make(chan int)
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			glog.Error(err)
			return err
		}
	case <-statusCh:
		glog.Infof("runner %s completed job %d, container id: %s", c.ID, job.Id, resp.ID)
		// c.RemoveContainer(job.Id)
		ch <- 1
	}
	<-ch
	return
}

func (c *Client) StopJob(id uint64) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		return
	}
	container_id := c.Containers[id].ShortID
	glog.Infof("runner %s stopping job %d, container id: %s", c.ID, id, container_id)
	if err := cli.ContainerStop(ctx, container_id, nil); err != nil {
		glog.Error(err)
		return err
	}
	glog.Infof("runner %s stopped job %d, container id: %s", c.ID, id, container_id)
	err = c.RemoveContainer(id)
	if err != nil {
		glog.Error(err)
		return err
	}
	return
}

func (c *Client) RemoveContainer(id uint64) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		return
	}
	// remove container
	container_id := c.Containers[id].ShortID
	glog.Infof("runner %s removing container id: %s, job id %d", c.ID, container_id, id)
	if err := cli.ContainerRemove(ctx, container_id, types.ContainerRemoveOptions{}); err != nil {
		glog.Error(err)
		return err
	}
	glog.Infof("runner %s removed container id: %s, job %d", c.ID, container_id, id)
	return
}
