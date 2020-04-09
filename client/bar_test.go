package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefreshBar(t *testing.T) {
	ct := createTestContainer()
	barPath := "testdata/bar.txt"
	ct.refreshBar(barPath)
	assert.Equal(t, uint32(499), ct.Job.CurrentStep)
	assert.Equal(t, uint32(500), ct.Job.TotalSteps)
}
