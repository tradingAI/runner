package plugins

import (
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

type Plugin struct{}

func (p *Plugin) GenerateCmds(input *pb.JobInput, id string) (cmds []string, err error) {
	return
}
