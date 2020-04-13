package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetLogicCPUNum(t *testing.T){
    assert.True(t, GetLogicCPUNum() > 0)
}


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

func TestGetMachineInfo(t *testing.T)  {
	cpuNum, totalMemory, availableMemeory, err := GetMachineInfo()
	assert.Nil(t, err)
	assert.True(t, cpuNum >= int32(1))
	assert.True(t, totalMemory >= int64(300 * 1024 * 1024))
	assert.True(t, availableMemeory >= int64(300 * 1024 * 1024))
}
