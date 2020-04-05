package client

import (
	"context"
	"os"
	"path"
	"bufio"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/golang/glog"
)

func (c *Client) createLogFile(id string) (logFilePath string) {
	if _, err := os.Stat(c.Conf.JobLogDir); os.IsNotExist(err) {
		err = os.MkdirAll(c.Conf.JobLogDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	logFileName := id + ".log"
	logFilePath = path.Join(c.Conf.JobLogDir, logFileName)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		glog.Fatal(err)
		return
	}
	f.Close()
	return
}

func writeLog(container_id string, filePath string) (err error) {
	ctx := context.Background()
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	reader, err := cli.ContainerLogs(ctx, container_id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if errClose := reader.Close(); errClose != nil {
			panic(errClose)
		}
	}()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		glog.Fatal(err)
		return
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		_, err = f.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			break
		}
	}
	if err != nil {
		panic(err)
	}

	return
}
