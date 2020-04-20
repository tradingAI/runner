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

type Plugin interface {
	GenerateCmds(input *pb.JobInput, id string) (cmds []string, err error)
	ParseBar(encode string) (currentStep uint32, totalSteps uint32, err error)
	ParseEval(encode string, jobId, modelId uint64) (out *pb.JobOutput, err error)
	ParseInfer(encode, date string, jobId, modelId uint64) (out *pb.JobOutput, err error)
}

func New(job *pb.Job) (p Plugin) {
	return NewTbasePlugin()
}
