package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConf(t *testing.T) {
	_, err := LoadConf()
	assert.Nil(t, err)
}
