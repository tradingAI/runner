package plugins

import (
	"fmt"

	mpb "github.com/tradingAI/proto/gen/go/model"
	pb "github.com/tradingAI/proto/gen/go/scheduler"
)

func GetTbaseInstallRepoCmds(repo string, tag string) (cmds []string) {
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
