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
		"git checkout -b v1.0.4 && pip install -e .",
		"cd /root/trade/tbase && git pull",
		"git checkout -b v0.1.5 && pip install -e .",
		"python -m tbase.run --alg ddpg",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateEvalCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseEvalJobInput()
	actual, _ := p.GenerateCmds(input)
	expected := []string{
		"cd /root/trade/tbase",
		"python -m tbase.runner --eval --bucket test_bucket --obj_name test_obj_name.tar.gz --eval_start 20190101 --eval_end 20200101",
	}
	assert.Equal(t, expected, actual)
}

func TestTbaseGenerateInferCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseInferJobInput()
	actual, _ := p.GenerateCmds(input)
	expected := []string{
		"cd /root/trade/tbase",
		"python -m tbase.runner --infer --bucket test_bucket --obj_name test_obj_name.tar.gz --infer_date 20200101",
	}
	assert.Equal(t, expected, actual)
}
