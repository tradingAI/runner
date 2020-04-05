package client

import (
	"context"
	"strconv"
	"bytes"

	"github.com/golang/glog"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
	"docker.io/go-docker/api/types/mount"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

const DEFAULT_IMAGE = "alpine"

// TODO
// const DEFAULT_IMAGE = "registry.cn-hangzhou.aliyuncs.com/tradingai/tbase:latest"

func getCmd(job *pb.Job) (cmd []string, err error) {
	// TODO
	return []string{"sh", "/root/test.sh"}, nil
}

func writeLog(id string)(err error){
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}
	out, err := cli.ContainerLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	outString := buf.String()
	glog.Info(outString)
	return nil
}

func (c *Client) CreateJob(job *pb.Job) (err error) {
	glog.Infof("runner %s creating job %d", c.ID, job.Id)
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	container_name := strconv.FormatUint(job.Id, 10)
	cmd, err := getCmd(job)
	if err != nil{
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: DEFAULT_IMAGE,
		Cmd:   cmd,
		Tty:   true,
	}, &container.HostConfig{
        Mounts: []mount.Mount{
            {
                Type:   mount.TypeBind,
                Source: "/Users/liuwen/Google/liulishuo/workspace/go/gopath/src/github.com/tradingAI/runner/test.sh",
                Target: "/root/test.sh",
            },
        },
    }, nil, container_name)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	c.Containers[job.Id] = Container{
		Name: strconv.FormatUint(job.Id, 10),
		ID: resp.ID,
		ShortID: resp.ID[:12],
		Job: job,
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
		case err := <-errCh:
			if err != nil {
				panic(err)
			}
		case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	outString := buf.String()
	glog.Info(outString)
	glog.Infof("runner %s created job %d, container id: %s", c.ID, job.Id, resp.ID)

	// write out into file, gorutine
	// TODO
	return
}

func (c *Client) StopJob(id uint64) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}
	// stop job
	container_id := c.Containers[id].ShortID
	glog.Infof("runner %s stopping job %d, container id: %s", c.ID, id, container_id)
	if err := cli.ContainerStop(ctx, container_id, nil); err != nil {
		panic(err)
	}
	glog.Infof("runner %s stopped job %d, container id: %s", c.ID, id, container_id)
	return
}


func (c *Client) RemoveContainer(id uint64) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}
	// remove container
	container_id := c.Containers[id].ShortID
	glog.Infof("runner %s removing container id: %s, job id %d", c.ID, container_id, id)
	if err := cli.ContainerRemove(ctx, container_id, types.ContainerRemoveOptions{}); err != nil {
		panic(err)
	}
	glog.Infof("runner %s removed container id: %s, job %d", c.ID, container_id, id)
	return
}
