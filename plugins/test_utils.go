package plugins

import(
    pb "github.com/tradingAI/proto/gen/go/scheduler"
    mpb "github.com/tradingAI/proto/gen/go/model"
)

func CreateTestTbaseTrainJobInput() (input *pb.JobInput){
    tenvsURL := "https://github.com/tradingAI/tenvs"
    tenvsTag := "v1.0.3"
    tenvsRepo := &mpb.InstallRepository{
        Repo: tenvsURL,
        Tag: tenvsTag,
        Cmds: GetTbaseInstallRepoCmds(tenvsURL, tenvsTag),
    }
    tbaseURL := "https://github.com/tradingAI/tbase"
    tbaseTag := "v0.1.2"
    tbaseRepo := &mpb.InstallRepository{
        Repo: tbaseURL,
        Tag: tbaseTag,
        Cmds: GetTbaseInstallRepoCmds(tbaseURL, tbaseTag),
    }
    parameters := make(map[string]string)
    parameters["alg"] = "ddpg"
    trainInput := &mpb.TbaseTrainInput{
        TenvsRepo: tenvsRepo,
        TbaseRepo: tbaseRepo,
        Parameters: parameters,
    }
    input = &pb.JobInput{
        Input: &pb.JobInput_TrainInput{trainInput},
    }
    return
}
