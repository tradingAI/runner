package client

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
	plugins "github.com/tradingAI/runner/plugins"
)

func (c *Client) getCmd(shellPath string) (cmd []string) {
	return []string{"sh", shellPath}
}

func (c *Client) createShellFile(job *pb.Job) (shellFilePath string) {
	if _, err := os.Stat(c.Conf.JobShellDir); os.IsNotExist(err) {
		err = os.MkdirAll(c.Conf.JobShellDir, 0755)
		if err != nil {
			glog.Error(err)
		}
	}
	shellFileName := strconv.FormatUint(job.Id, 10) + ".sh"
	shellFilePath = path.Join(c.Conf.JobShellDir, shellFileName)
	f, err := os.OpenFile(shellFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		glog.Error(err)
	}
	// TODO: write cmds
	p := &plugins.TbasePlugin{}
	commands, err := p.GenerateCmds(job.Input)
	glog.Infof("GenerateCmds len %d", cap(commands))
	if err != nil {
		glog.Error(err)
	}
	for _, cmd := range commands {
		line := fmt.Sprintf("%s\n", cmd)
		glog.Info(line)
		f.Write([]byte(line))
	}
	f.Close()
	return
}
