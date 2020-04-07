package plugins

import (
	"errors"
	"fmt"
	"sort"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

const DEFAULT_TBASE_MODEL_DIR = "/root/data/model/"
const DEFAULT_TBASE_PROGRESS_BAR_PATH = "/root/progress_bar/bar.txt"
const DEFAULT_TBASE_TENSORBOARD_DIR = "/root/tensorboard/"
const DEFAULT_TBASE_INFER_PATH = "/root/inferences/infer.txt"
const DEFAULT_TBASE_EVAL_PATH = "/root/evals/eval.txt"

type TbasePlugin struct{}

func (p *TbasePlugin) GenerateCmds(input *pb.JobInput) (cmds []string, err error) {
	switch input.GetInput().(type) {
	case *pb.JobInput_EvalInput:
		glog.Info("JobInput_EvalInput")
		cmds, err = p.getEvalJobCmds(input)
		if err != nil {
			glog.Error(err)
			return
		}
	case *pb.JobInput_InferInput:
		glog.Info("JobInput_InferInput")
		cmds, err = p.getInferJobCmds(input)
		if err != nil {
			glog.Error(err)
			return
		}
	case *pb.JobInput_TrainInput:
		glog.Info("JobInput_TrainInput")
		cmds, err = p.getTrainJobCmds(input)
		if err != nil {
			glog.Error(err)
			return
		}
	default:
		err = errors.New("plugins GenerateCmds input invalid")
		glog.Error(err)
		return
	}
	return
}

func (p *TbasePlugin) getTrainJobCmds(input *pb.JobInput) (cmds []string, err error) {
	// https://github.com/tradingAI/tbase/blob/2ccac243409fe93c15c0ceb4cff9fe419166590e/Dockerfile
	// tenvs
	tenvsTag := input.GetTrainInput().GetTenvsTag()
	installTenvsCmds := GetTbaseInstallRepoCmds("tenvs", tenvsTag)
	cmds = append(cmds, installTenvsCmds...)
	// tbase
	tbaseTag := input.GetTrainInput().GetTbaseTag()
	installTbaseCmds := GetTbaseInstallRepoCmds("tbase", tbaseTag)
	cmds = append(cmds, installTbaseCmds...)
	// run commands
	parametersStr := ""
	parameters := input.GetTrainInput().GetParameters()
	parameters["model_dir"] = DEFAULT_TBASE_MODEL_DIR
	parameters["progress_bar_path"] = DEFAULT_TBASE_PROGRESS_BAR_PATH
	parameters["tensorboard_dir"] = DEFAULT_TBASE_TENSORBOARD_DIR
	sorted_keys := make([]string, 0)
	for k, _ := range parameters {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	for _, k := range sorted_keys {
		parametersStr = fmt.Sprintf("%s --%s %s", parametersStr, k, parameters[k])
	}
	runCmd := fmt.Sprintf("python -m trunner.tbase%s", parametersStr)
	cmds = append(cmds, runCmd)
	return
}

func (p *TbasePlugin) getEvalJobCmds(input *pb.JobInput) (cmds []string, err error) {
	start := input.GetEvalInput().GetStart()
	end := input.GetEvalInput().GetEnd()
	// https://github.com/tradingAI/tbase/blob/2ccac243409fe93c15c0ceb4cff9fe419166590e/Dockerfile
	cmds = append(cmds, "cd /root/trade/tbase")
	// tbase会读取model中的meta 版本信息，自动checkout到相应版本运行程序
	runCmd := fmt.Sprintf("python -m trunner.tbase --eval --model_dir %s --eval_result_path %s --eval_start %s --eval_end %s",
		DEFAULT_TBASE_MODEL_DIR, DEFAULT_TBASE_EVAL_PATH, start, end)
	cmds = append(cmds, runCmd)
	return
}

func (p *TbasePlugin) getInferJobCmds(input *pb.JobInput) (cmds []string, err error) {
	inferDate := input.GetInferInput().GetDate()
	// https://github.com/tradingAI/tbase/blob/21a72ee53b7b7c2c1a976d8e1c2a6d858de64564/Dockerfile#L12
	cmds = append(cmds, "cd /root/trade/tbase")
	// tbase会读取model中的meta 版本信息，自动checkout到相应版本运行程序
	runCmd := fmt.Sprintf("python -m trunner.tbase --infer --model_dir %s --infer_result_path %s --infer_date %s",
		DEFAULT_TBASE_MODEL_DIR, DEFAULT_TBASE_INFER_PATH, inferDate)
	cmds = append(cmds, runCmd)
	return
}
