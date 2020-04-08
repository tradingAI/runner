package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstallTbaseRepoCmds(t *testing.T) {
	tenvsURL := "tenvs"
	tenvsTag := "v1.0.4"
	actual := GetTbaseInstallRepoCmds(tenvsURL, tenvsTag)
	expected := []string{
		"cd /root/trade/tenvs && git pull",
		"git checkout tags/v1.0.4 -b v1.0.4-branch && pip install -e .",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateTrainCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseTrainJobInput()
	actual, _ := p.GenerateCmds(input, "0")
	runCmd := fmt.Sprintf("python -m trunner.tbase --alg ddpg --model_dir %smodels --progress_bar_path %sprogress_bars/0 --tensorboard_dir %stensorboards",
		ROOT_DATA_DIR, ROOT_DATA_DIR, ROOT_DATA_DIR)
	expected := []string{
		"cd /root/trade/tenvs && git pull",
		"git checkout tags/v1.0.8 -b v1.0.8-branch && pip install -e .",
		"cd /root/trade/tbase && git pull",
		"git checkout tags/v0.1.8 -b v0.1.8-branch && pip install -e .",
		runCmd,
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateEvalCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseEvalJobInput()
	actual, _ := p.GenerateCmds(input, "0")
	runCmd := fmt.Sprintf("python -m trunner.tbase --eval --model_dir %smodels/0 --eval_result_path %sevals/0 --eval_start 20190101 --eval_end 20200101",
		ROOT_DATA_DIR, ROOT_DATA_DIR)
	expected := []string{
		"cd /root/trade/tbase",
		runCmd,
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateInferCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseInferJobInput()
	actual, _ := p.GenerateCmds(input, "0")
	runCmd := fmt.Sprintf("python -m trunner.tbase --infer --model_dir %smodels/0 --infer_result_path %sinfers/0 --infer_date 20200101",
		ROOT_DATA_DIR, ROOT_DATA_DIR)
	expected := []string{
		"cd /root/trade/tbase",
		runCmd,
	}
	assert.Equal(t, expected, actual)
}
