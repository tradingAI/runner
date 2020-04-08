package plugins

import (
	"os"
	"path"

	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

var ROOT_DATA_DIR = os.Getenv("RUNNER_DATA_DIR")
var JOB_SHELL_DIR = path.Join(ROOT_DATA_DIR, "shells")
var JOB_LOG_DIR = path.Join(ROOT_DATA_DIR, "logs")
var MODEL_DIR = path.Join(ROOT_DATA_DIR, "models")
var PROGRESS_BAR_DIR = path.Join(ROOT_DATA_DIR, "progress_bars")
var TENSORBOARD_DIR = path.Join(ROOT_DATA_DIR, "tensorboards")
var INFER_DIR = path.Join(ROOT_DATA_DIR, "infers")
var EVAL_DIR = path.Join(ROOT_DATA_DIR, "evals")

type Plugin struct{}

func (p *Plugin) GenerateCmds(input *pb.JobInput, id string) (cmds []string, err error) {
	return
}
