package runner

import (
	"bufio"
	"context"
	"os"
	"path"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/golang/glog"
)

func (r *Runner) createLogFile(id string) (logFilePath string) {
	if _, err := os.Stat(r.Conf.JobLogDir); os.IsNotExist(err) {
		err = os.MkdirAll(r.Conf.JobLogDir, 0755)
		if err != nil {
			glog.Error(err)
		}
	}
	logFilePath = path.Join(r.Conf.JobLogDir, id)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		glog.Error(err)
		return
	}
	f.Close()
	return
}

func writeLog(containerId string, filePath string) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		glog.Error(err)
		return
	}
	defer cli.Close()

	reader, err := cli.ContainerLogs(ctx, containerId, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
	})
	if err != nil {
		glog.Error(err)
		return
	}
	defer func() {
		if errClose := reader.Close(); errClose != nil {
			glog.Error(errClose)
			return
		}
	}()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		glog.Error(err)
		return
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		_, err = f.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			glog.Error(err)
			return
		}
	}
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
