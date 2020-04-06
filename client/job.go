package client

import (
	"context"
	"strconv"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/mount"
	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func (c *Client) CreateJob(job *pb.Job) (err error) {
	glog.Infof("runner %s creating job %d", c.ID, job.Id)
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
	}

	shellFilePath := c.createShellFile(job)
	jobIdStr := strconv.FormatUint(job.Id, 10)
	logFilePath := c.createLogFile(jobIdStr)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        DEFAULT_IMAGE,
		Cmd:          c.getCmd(TARGET_SHELL_PATH),
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: shellFilePath,
				Target: TARGET_SHELL_PATH,
			},
		},
	}, nil, jobIdStr)
	if err != nil {
		glog.Error(err)
	}

	glog.Infof("runner %s created job %d, container id: %s", c.ID, job.Id, resp.ID)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		glog.Error(err)
	}

	c.Containers[job.Id] = Container{
		Name:    strconv.FormatUint(job.Id, 10),
		ID:      resp.ID,
		ShortID: resp.ID[:12],
		Job:     job,
	}

	// stream read container log and write into file
	go writeLog(resp.ID, logFilePath)

	ch := make(chan int)
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			glog.Error(err)
		}
	case <-statusCh:
		glog.Infof("runner %s completed job %d, container id: %s", c.ID, job.Id, resp.ID)
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
	}
	container_id := c.Containers[id].ShortID
	glog.Infof("runner %s stopping job %d, container id: %s", c.ID, id, container_id)
	if err := cli.ContainerStop(ctx, container_id, nil); err != nil {
		glog.Error(err)
	}
	glog.Infof("runner %s stopped job %d, container id: %s", c.ID, id, container_id)
	return
}

func (c *Client) RemoveContainer(id uint64) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
	}
	// remove container
	container_id := c.Containers[id].ShortID
	glog.Infof("runner %s removing container id: %s, job id %d", c.ID, container_id, id)
	if err := cli.ContainerRemove(ctx, container_id, types.ContainerRemoveOptions{}); err != nil {
		glog.Error(err)
	}
	glog.Infof("runner %s removed container id: %s, job %d", c.ID, container_id, id)
	return
}
