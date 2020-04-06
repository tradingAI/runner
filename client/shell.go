package client

import (
	"os"
	"path"
	"strconv"
	"fmt"

	"github.com/golang/glog"
	plugins "github.com/tradingAI/runner/plugins"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
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
		return
	}
	// TODO: write cmds
	p := &plugins.TbasePlugin{}
	commands, err := p.GenerateCmds(job.Input)
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
