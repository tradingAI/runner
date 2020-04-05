package client

import (
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func (c *Client) getCmd(shellPath string) (cmd []string) {
	return []string{"sh", shellPath}
}

func (c *Client) createShellFile(job *pb.Job) (shellFilePath string) {
	if _, err := os.Stat(c.Conf.JobShellDir); os.IsNotExist(err) {
		err = os.MkdirAll(c.Conf.JobShellDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	shellFileName := strconv.FormatUint(job.Id, 10) + ".sh"
	shellFilePath = path.Join(c.Conf.JobShellDir, shellFileName)
	f, err := os.OpenFile(shellFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		glog.Fatal(err)
		return
	}
	// TODO: write cmds
	f.Write([]byte("echo " + strings.Repeat("==", 40) + "\n"))
	for i := 0; i < 10; i++ {
		f.Write([]byte("echo " + strconv.Itoa(i) + "\n"))
		f.Write([]byte("echo " + strings.Repeat("====", i + 1) + "\n"))
		f.Write([]byte("sleep 1s\n"))
		f.Write([]byte("date\n"))
	}

	f.Close()
	return
}
