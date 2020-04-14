package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetMemory(t *testing.T){
    totalMemory, _, err := GetMemory()
    // NOTE(wen): 默认内存大于 300M
    byte100M := int64(300 * 1024 * 1024)
    assert.True(t, totalMemory > byte100M)
    assert.Nil(t, err)
}

func TestGetPhysicalCPUNum(t *testing.T){
    actual, err := GetPhysicalCPUNum()
    assert.Nil(t, err)
    assert.True(t, actual >= int32(1))
}

func TestUpdateMemory(t *testing.T){
	m, err := NewMachine()
	assert.Nil(t, err)
	m.UpdateMemory()
	assert.True(t, m.AvailableMemory > 0)
}

func TestUpdateCPUUtilization(t *testing.T){
	m, err := NewMachine()
	assert.Equal(t, float32(0), m.CPUUtilization)
	assert.Nil(t, err)
	m.UpdateCPUUtilization()
	assert.True(t, m.CPUUtilization > 0)
}

// TODO: add gpu test in GPU evironment
