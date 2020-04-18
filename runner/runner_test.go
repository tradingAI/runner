package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := creatTestRunner()
	assert.Equal(t, 36, len(c.ID))
	assert.Equal(t, "test_token", c.Conf.Token)
}
