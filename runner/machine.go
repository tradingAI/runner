package runner

import (
	"runtime"

	"github.com/golang/glog"
	"github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
)

// NOTE: 这里获取的机器信息都是容机内的虚拟机的信息, linux系统可以通过
// docker run -it -v /proc:/hostinfo/proc:ro 镜像名
// 并设置环境变量 HOST_PROC=/hostinfo/proc 来获取宿主机器的信息

func GetLogicCPUNum() (n int32) {
	cpuNum := runtime.GOMAXPROCS(runtime.NumCPU())
	return int32(cpuNum)
}

// 获取总内存，可用内存信息
func GetMemory() (totalMemory, availableMemeory int64, err error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		glog.Error(err)
		return
	}
	totalMemory = int64(v.Total)
	availableMemeory = int64(v.Available)
	return
}

// 获取机器的cpu数量
func GetPhysicalCPUNum() (n int32, err error){
    physicalCnt, err := cpu.Counts(false)
    if err != nil {
		glog.Error(err)
		return
	}
    n = int32(physicalCnt)
    return
}

func GetMachineInfo()(cpuNum int32, totalMemory, availableMemeory int64, err error){
	cpuNum, err = GetPhysicalCPUNum()
	if err != nil {
		glog.Error(err)
		return
	}
	totalMemory, availableMemeory, err = GetMemory()
    return
}
