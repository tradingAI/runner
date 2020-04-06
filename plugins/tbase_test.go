package plugins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstallTbaseRepoCmds(t *testing.T) {
	tenvsURL := "https://github.com/tradingAI/tenvs"
	tenvsTag := "v1.0.3"
	actual := GetTbaseInstallRepoCmds(tenvsURL, tenvsTag)
	expected := []string{
		"git clone https://github.com/tradingAI/tenvs.git",
		"cd tenvs",
		"git checkout -b v1.0.3",
		"pip install -e .",
	}
	assert.Equal(t, expected, actual)
}

func TestGenerateCmds(t *testing.T) {
	p := &TbasePlugin{}
	input := CreateTestTbaseTrainJobInput()
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
