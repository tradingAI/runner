package plugins

import(
    pb "github.com/tradingAI/proto/gen/go/scheduler"
    "github.com/golang/glog"
    "fmt"
)

type TbasePlugin struct {}

func (p *TbasePlugin) GenerateCmds(input *pb.JobInput)(cmds []string, err error){
    switch input.GetInput().(type) {
    case *pb.JobInput_EvalInput:
        // TODO
        glog.Info("JobInput_EvalInput")
    case *pb.JobInput_InferInput:
        // TODO
        glog.Info("JobInput_InferInput")
    case *pb.JobInput_TrainInput:
        // TODO
        glog.Info("JobInput_TrainInput")
        return p.getTrainJobCmds(input)
    }
    return
}

func (p *TbasePlugin) getTrainJobCmds(input *pb.JobInput)(cmds []string, err error){
    // add run dir
    cmds = append(cmds, "mkdir -p /root/runner/")
    cmds = append(cmds, "cd /root/runner/")
    // install tenvs dir
    cmds = append(cmds, input.GetTrainInput().GetTenvsRepo().GetCmds()...)
    // tbase
    cmds = append(cmds, "cd /root/runner/")
    cmds = append(cmds, input.GetTrainInput().GetTbaseRepo().GetCmds()...)
    // run commands
    // TODO: 根据parameters参数构造run命令
    runCmd := fmt.Sprintf("python -m tbase.run --%s %s", "alg", input.GetTrainInput().GetParameters()["alg"])
    cmds = append(cmds, runCmd)
    return
}
