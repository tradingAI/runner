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
	installCmd := fmt.Sprintf("git checkout -b %s && pip install -e .", tag)
	cmds = []string{pullCmd, installCmd}
	return
}

func CreateDefaultTbaseTrainJobInput() (input *pb.JobInput) {
	parameters := make(map[string]string)
	parameters["alg"] = "ddpg"
	trainInput := &mpb.TbaseTrainInput{
		TenvsTag:   "v1.0.4",
		TbaseTag:   "v0.1.5",
		Parameters: parameters,
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
			Id:          uint64(123456789),
			Name:        "ddpg",
			Version:     "v1.0.0",
			Description: "",
			FileType:    "tar.gz",
			User:        &cpb.User{Id: uint64(1)},
			Status:      cpb.ModelStatus_SUCCESS,
			Bucket:      "test_bucket",
			ObjName:     "test_obj_name.tar.gz",
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
			Id:          uint64(123456789),
			Name:        "ddpg",
			Version:     "v1.0.0",
			Description: "",
			FileType:    "tar.gz",
			User:        &cpb.User{Id: uint64(1)},
			Status:      cpb.ModelStatus_SUCCESS,
			Bucket:      "test_bucket",
			ObjName:     "test_obj_name.tar.gz",
		},
		Date: "20200101",
	}
	input = &pb.JobInput{
		Input: &pb.JobInput_InferInput{infeInput},
	}
	return
}
