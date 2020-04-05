package client

import (
	// "bytes"
	"context"
	"io"
	"os"
	"path"
	"time"

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
	out, err := cli.ContainerLogs(ctx, container_id, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		glog.Fatal(err)
		return
	}
	buf := make([]byte, 8)
	for {
		n, err := out.Read(buf)
		if err == io.EOF {
			time.Sleep(1 * time.Second)
			break
		}
		outString := string(buf[:n])
		f.Write([]byte(outString))
	}
	f.Close()
	return
}
