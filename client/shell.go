package client

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	"github.com/tradingAI/runner/plugins"
)

func (c *Client) getCmd(shellPath string) (cmd []string) {
	return []string{"sh", shellPath}
}

func (c *Client) createShellFile(job *pb.Job, p plugins.Plugin) (err error) {
	if _, err := os.Stat(c.Conf.JobShellDir); os.IsNotExist(err) {
		err = os.MkdirAll(c.Conf.JobShellDir, 0755)
		if err != nil {
			glog.Error(err)
			return err
		}
	}
	id := strconv.FormatUint(job.Id, 10)
	shellFilePath := path.Join(c.Conf.JobShellDir, id)
	f, err := os.OpenFile(shellFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		glog.Error(err)
		return err
	}
	commands, err := p.GenerateCmds(job.Input, id)
	glog.Infof("GenerateCmds len %d", cap(commands))
	if err != nil {
		glog.Error(err)
		return
	}
	for _, cmd := range commands {
		line := fmt.Sprintf("%s\n", cmd)
		f.Write([]byte(line))
	}
	f.Close()
	return
}
