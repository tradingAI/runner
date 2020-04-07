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

func TestGenerateCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateDefaultTbaseTrainJobInput()
	actual, _ := p.GenerateCmds(input, nil)
	expected := []string{
		"cd /root/trade/tenvs && git pull",
		"git checkout -b v1.0.4 && pip install -e .",
		"cd /root/trade/tbase && git pull",
		"git checkout -b v0.1.5 && pip install -e .",
		"python -m tbase.run --alg ddpg",
	}
	assert.Equal(t, expected, actual)
}
