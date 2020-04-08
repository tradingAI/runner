package plugins

import (
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

const MODEL_DIR = "/root/data/model/"
const PROGRESS_BAR_DIR = "/root/data/progress_bar/"
const TENSORBOARD_DIR = "/root/data/tensorboard/"
const INFER_DIR = "/root/data/inferences/"
const EVAL_DIR = "/root/data/evals/"

type Plugin struct{}

func (p *Plugin) GenerateCmds(input *pb.JobInput, id string) (cmds []string, err error) {
	return
}
