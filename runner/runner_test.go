package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := creatTestRunner()
	assert.Equal(t, "test_runner_id", c.ID)
	assert.Equal(t, "test_token", c.Token)
}
