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
	actual, _ := p.GenerateCmds(input)
	expected := []string{
		"cd /root/trade/tenvs && git pull",
		"git checkout -b v1.0.5 && pip install -e .",
		"cd /root/trade/tbase && git pull",
		"git checkout -b v0.1.6 && pip install -e .",
		"python -m trunner.tbase --alg ddpg --model_dir /root/data/model/ --progress_bar_path /root/progress_bar/bar.txt --tensorboard_dir /root/tensorboard/",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateEvalCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseEvalJobInput()
	actual, _ := p.GenerateCmds(input)
	expected := []string{
		"cd /root/trade/tbase",
		"python -m trunner.tbase --eval --model_dir /root/data/model/ --eval_result_path /root/evals/eval.txt --eval_start 20190101 --eval_end 20200101",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateInferCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseInferJobInput()
	actual, _ := p.GenerateCmds(input)
	expected := []string{
		"cd /root/trade/tbase",
		"python -m trunner.tbase --infer --model_dir /root/data/model/ --infer_result_path /root/inferences/infer.txt --infer_date 20200101",
	}
	assert.Equal(t, expected, actual)
}
