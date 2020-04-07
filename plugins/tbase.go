package plugins

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/minio/minio-go/v6"
	minio2 "github.com/tradingAI/go/s3/minio"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

type TbasePlugin struct{}

func (p *TbasePlugin) GenerateCmds(input *pb.JobInput, minioClient *minio.Client) (cmds []string, err error) {
	switch input.GetInput().(type) {
	case *pb.JobInput_EvalInput:
		glog.Info("JobInput_EvalInput")
		return p.getEvalJobCmds(input, minioClient)
	case *pb.JobInput_InferInput:
		// TODO
		glog.Info("JobInput_InferInput")
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

func (p *TbasePlugin) getEvalJobCmds(input *pb.JobInput, minioClient *minio.Client) (cmds []string, err error) {
	// download model
	// TODO: change after https://github.com/tradingAI/proto/pull/16
	modelBucket := input.GetEvalInput().GetModel().GetName()
	objName := input.GetEvalInput().GetModel().GetName()
	fp := "/root/runner/model.tar.gz"
	err = minio2.MinioDownload(minioClient, modelBucket, fp, objName)
	if err != nil {
		glog.Error(err)
	}
	model_dir := "/root/runner/model"
	// TODO: 解压模型到 model_dir
	start := input.GetEvalInput().GetStart()
	end := input.GetEvalInput().GetEnd()
	// https://github.com/tradingAI/tbase/blob/21a72ee53b7b7c2c1a976d8e1c2a6d858de64564/Dockerfile#L12
	cmds = append(cmds, "cd /root/trade/tbase")
	// tbase会读取model中的meta 版本信息，自动checkout到相应版本运行程序
	runCmd := fmt.Sprintf("python -m tbase.runner --eval --model_dir %s --eval_start %s --eval_end %s", model_dir, start, end)
	cmds = append(cmds, runCmd)
	return
}

func (p *TbasePlugin) getInferJobCmds(input *pb.JobInput, minioClient *minio.Client) (cmds []string, err error) {
	// download model
	// TODO: change after https://github.com/tradingAI/proto/pull/16
	modelBucket := input.GetInferInput().GetModel().GetName()
	objName := input.GetInferInput().GetModel().GetName()
	fp := "/root/runner/model.tar.gz"
	err = minio2.MinioDownload(minioClient, modelBucket, fp, objName)
	if err != nil {
		glog.Error(err)
	}
	model_dir := "/root/runner/model"
	// TODO: 解压模型到 model_dir
	inferDate := input.GetInferInput().GetDate()
	// https://github.com/tradingAI/tbase/blob/21a72ee53b7b7c2c1a976d8e1c2a6d858de64564/Dockerfile#L12
	cmds = append(cmds, "cd /root/trade/tbase")
	// tbase会读取model中的meta 版本信息，自动checkout到相应版本运行程序
	runCmd := fmt.Sprintf("python -m tbase.runner --infer --model_dir %s --infer_date %s", model_dir, inferDate)
	cmds = append(cmds, runCmd)
	return
}
