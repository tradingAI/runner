package plugins

import (
	"fmt"

	cpb "github.com/tradingAI/proto/gen/go/common"
	mpb "github.com/tradingAI/proto/gen/go/model"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func GetTbaseInstallRepoCmds(repo string, tag string) (cmds []string) {
	// https://github.com/tradingAI/tbase/blob/2ccac243409fe93c15c0ceb4cff9fe419166590e/Dockerfile
	pullCmd := fmt.Sprintf("cd /root/trade/%s && git pull", repo)
	installCmd := fmt.Sprintf("git checkout tags/%s -b %s-branch && pip install -e .", tag, tag)
	cmds = []string{pullCmd, installCmd}
	return
}

func CreateDefaultTbaseTrainJobInput() (input *pb.JobInput) {
	parameters := make(map[string]string)
	parameters["alg"] = "ddpg"
	parameters["max_iter_num"] = "10"
	parameters["warm_up"] = "1000"
	parameters["seed"] = "0"

	trainInput := &mpb.TbaseTrainInput{
		TenvsTag:   "v1.0.8",
		TbaseTag:   "v0.1.8",
		Parameters: parameters,
		Bucket: "tbase",
		ModelFileDir: "test_user/model/",
		TensorboardFileDir: "test_user/tensorboard/",
	}
	input = &pb.JobInput{
		Input: &pb.JobInput_TrainInput{trainInput},
	}
	return
}

func CreateDefaultTbaseEvalJobInput() (input *pb.JobInput) {
	parameters := make(map[string]string)
	parameters["alg"] = "ddpg"
	evalInput := &mpb.TbaseEvaluateInput{
		Model: &cpb.Model{
			Id:          uint64(2),
			Name:        "ddpg",
			Version:     "v1.0.0",
			Description: "",
			FileType:    "application/zip",
			User:        &cpb.User{Id: uint64(1)},
			Status:      cpb.ModelStatus_SUCCESS,
			Bucket:      "tbase",
			ObjName:     "/test_user/model/22222.zip",
		},
		Start: "20190101",
		End:   "20200101",
	}
	input = &pb.JobInput{
		Input: &pb.JobInput_EvalInput{evalInput},
	}
	return
}

func CreateDefaultTbaseInferJobInput() (input *pb.JobInput) {
	parameters := make(map[string]string)
	parameters["alg"] = "ddpg"
	infeInput := &mpb.TbaseInferInput{
		Model: &cpb.Model{
			Id:          uint64(1),
			Name:        "ddpg",
			Version:     "v1.0.0",
			Description: "",
			FileType:    "application/zip",
			User:        &cpb.User{Id: uint64(1)},
			Status:      cpb.ModelStatus_SUCCESS,
			Bucket:      "tbase",
			ObjName:     "/test_user/model/22222.zip",
		},
		Date: "20200101",
	}
	input = &pb.JobInput{
		Input: &pb.JobInput_InferInput{infeInput},
	}
	return
}
