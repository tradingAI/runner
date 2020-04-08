package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstallTbaseRepoCmds(t *testing.T) {
	tenvsURL := "tenvs"
	tenvsTag := "v1.0.4"
	actual := GetTbaseInstallRepoCmds(tenvsURL, tenvsTag)
	expected := []string{
		"cd /root/trade/tenvs && git pull",
		"git checkout -b v1.0.4 && pip install -e .",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateTrainCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseTrainJobInput()
	actual, _ := p.GenerateCmds(input, "0")
	expected := []string{
		"cd /root/trade/tenvs && git pull",
		"git checkout -b v1.0.8 && pip install -e .",
		"cd /root/trade/tbase && git pull",
		"git checkout -b v0.1.8 && pip install -e .",
		"python -m trunner.tbase --alg ddpg --model_dir /root/data/model/ --progress_bar_path /root/data/progress_bar/0 --tensorboard_dir /root/data/tensorboard/",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateEvalCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseEvalJobInput()
	actual, _ := p.GenerateCmds(input, "0")
	expected := []string{
		"cd /root/trade/tbase",
		"python -m trunner.tbase --eval --model_dir /root/data/model/0 --eval_result_path /root/data/evals/0 --eval_start 20190101 --eval_end 20200101",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateInferCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseInferJobInput()
	actual, _ := p.GenerateCmds(input, "0")
	expected := []string{
		"cd /root/trade/tbase",
		"python -m trunner.tbase --infer --model_dir /root/data/model/0 --infer_result_path /root/data/inferences/0 --infer_date 20200101",
	}
	assert.Equal(t, expected, actual)
}
