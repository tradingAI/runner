package plugins

import (
	"fmt"

	"github.com/golang/glog"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

type TbasePlugin struct{}

func (p *TbasePlugin) GenerateCmds(input *pb.JobInput) (cmds []string, err error) {
	switch input.GetInput().(type) {
	case *pb.JobInput_EvalInput:
		glog.Info("JobInput_EvalInput")
		return p.getEvalJobCmds(input)
	case *pb.JobInput_InferInput:
		glog.Info("JobInput_InferInput")
		return p.getInferJobCmds(input)
	case *pb.JobInput_TrainInput:
		glog.Info("JobInput_TrainInput")
		return p.getTrainJobCmds(input)
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
	for k, v := range input.GetTrainInput().GetParameters() {
		parametersStr = fmt.Sprintf("%s --%s %s", parametersStr, k, v)
	}
	runCmd := fmt.Sprintf("python -m tbase.run%s", parametersStr)
	cmds = append(cmds, runCmd)
	return
}

func (p *TbasePlugin) getEvalJobCmds(input *pb.JobInput) (cmds []string, err error) {
	start := input.GetEvalInput().GetStart()
	end := input.GetEvalInput().GetEnd()
	bucket := input.GetEvalInput().GetModel().GetBucket()
	objName := input.GetEvalInput().GetModel().GetObjName()
	// https://github.com/tradingAI/tbase/blob/2ccac243409fe93c15c0ceb4cff9fe419166590e/Dockerfile
	cmds = append(cmds, "cd /root/trade/tbase")
	// tbase会读取model中的meta 版本信息，自动checkout到相应版本运行程序
	runCmd := fmt.Sprintf("python -m tbase.runner --eval --bucket %s --obj_name %s --eval_start %s --eval_end %s", bucket, objName, start, end)
	cmds = append(cmds, runCmd)
	return
}

func (p *TbasePlugin) getInferJobCmds(input *pb.JobInput) (cmds []string, err error) {
	inferDate := input.GetInferInput().GetDate()
	bucket := input.GetInferInput().GetModel().GetBucket()
	objName := input.GetInferInput().GetModel().GetObjName()
	// https://github.com/tradingAI/tbase/blob/21a72ee53b7b7c2c1a976d8e1c2a6d858de64564/Dockerfile#L12
	cmds = append(cmds, "cd /root/trade/tbase")
	// tbase会读取model中的meta 版本信息，自动checkout到相应版本运行程序
	runCmd := fmt.Sprintf("python -m tbase.runner --infer --bucket %s --obj_name %s --infer_date %s", bucket, objName, inferDate)
	cmds = append(cmds, runCmd)
	return
}
