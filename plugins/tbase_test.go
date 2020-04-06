package plugins

import(
    "testing"
    "fmt"
    "strings"

    "github.com/stretchr/testify/assert"

    pb "github.com/tradingAI/proto/gen/go/scheduler"
    mpb "github.com/tradingAI/proto/gen/go/model"
)


func getCmds(url string, tag string) (cmds []string){
    segments := strings.Split(url, "/")
    name := segments[cap(segments) - 1]
    cloneCmd := fmt.Sprintf("git clone %s.git", url)
    checkoutCmd := fmt.Sprintf("git checkout -b %s", tag)
    cdCmd := fmt.Sprintf("cd %s", name)
    pipCmd := "pip install -e ."
    cmds = []string{cloneCmd, cdCmd, checkoutCmd, pipCmd}
    return
}

func TestGenerateCmds(t *testing.T){
    p := &TbasePlugin{}
    tenvsURL := "https://github.com/tradingAI/tenvs"
    tenvsTag := "v1.0.3"
    tenvsRepo := &mpb.InstallRepository{
        Repo: tenvsURL,
        Tag: tenvsTag,
        Cmds: getCmds(tenvsURL, tenvsTag),
    }
    tbaseURL := "https://github.com/tradingAI/tbase"
    tbaseTag := "v0.1.2"
    tbaseRepo := &mpb.InstallRepository{
        Repo: tbaseURL,
        Tag: tbaseTag,
        Cmds: getCmds(tbaseURL, tbaseTag),
    }
    parameters := make(map[string]string)
    parameters["alg"] = "ddpg"
    trainInput := &mpb.TbaseTrainInput{
        TenvsRepo: tenvsRepo,
        TbaseRepo: tbaseRepo,
        Parameters: parameters,
    }
    input := &pb.JobInput{
        Input: &pb.JobInput_TrainInput{trainInput},
    }
    actual, _ := p.GenerateCmds(input)
    expected := []string{
        "mkdir -p /root/runner/",
        "cd /root/runner/",
        "git clone https://github.com/tradingAI/tenvs.git",
        "cd tenvs",
        "git checkout -b v1.0.3",
        "pip install -e .",
        "cd /root/runner/",
        "git clone https://github.com/tradingAI/tbase.git",
        "cd tbase",
        "git checkout -b v0.1.2",
        "pip install -e .",
        "python -m tbase.run --alg ddpg",
    }
    assert.Equal(t, expected, actual)
}
