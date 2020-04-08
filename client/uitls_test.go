package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustCreatePath(t *testing.T) {
    err := mustCreatePath("/tmp/a/b/c/123456789")
	assert.Nil(t, err)
}

func TestMustCreateDir(t *testing.T) {
    err := mustCreateDir("/tmp/a/b/c/d/123")
	assert.Nil(t, err)
}
