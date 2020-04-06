package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	conf, _ := LoadConf()
	c, _ := New(conf)
	assert.Equal(t, "test_runner_id", c.ID)
	assert.Equal(t, "test_token", c.Token)
}
