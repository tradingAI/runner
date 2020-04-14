package runner

import (
	"time"

	"github.com/tradingAI/gpu-monitoring-tools/bindings/go/nvml"
	"github.com/golang/glog"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// NOTE: 这里获取的机器信息都是容机内的虚拟机的信息, linux系统可以通过
// docker run -it -v /proc:/hostinfo/proc:ro 镜像名
// 并设置环境变量 HOST_PROC=/hostinfo/proc 来获取宿主机器的信息

type Machine struct {
	gpuAvailable       bool
	GPUNum             int32
	GPUsIndex          []int32
	GPUMemory          int64
	AvailableGPUMemory int64
	GPUDevices         []*nvml.Device
	CPUNum             int32
	CPUUtilization     float32
	Memory             int64
	AvailableMemory    int64
}

func isGPUAvailable() (gpuAvailable bool, err error) {
	// TODO: 当GPU可用时
	err = nvml.Init()
	if err != nil {
		glog.Error(err)
		if err.Error() == "could not load NVML library" {
			glog.Info("isGPUAvailable: false, could not load NVML library")
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewMachine() (m *Machine, err error) {
	gpuAvailable, err := isGPUAvailable()
	if err != nil {
		glog.Error(err)
		return
	}
	var gpuNum uint
	var gpusIndex []int32
	var totalGPUMemory int64
	var devices []*nvml.Device
	if gpuAvailable {
		gpuNum, err = nvml.GetDeviceCount()
		if err != nil {
			glog.Error(err)
			return
		}
		for i := uint(0); i < gpuNum; i++ {
			device, err := nvml.NewDevice(i)
			if err != nil {
				glog.Error(err)
				return m, err
			}
			devices = append(devices, device)
			gpusIndex = append(gpusIndex, int32(i))
		}
		totalGPUMemory = int64(0)
		for _, dev := range devices {
			if mem := dev.Memory; mem != nil {
				totalGPUMemory += int64(*mem)
			}
		}
	}

	cpuNum, err := GetPhysicalCPUNum()
	if err != nil {
		glog.Error(err)
		return
	}
	totalMemory, availableMemeory, err := GetMemory()
	if err != nil {
		glog.Error(err)
		return
	}
	m = &Machine{
		GPUNum:   int32(gpuNum),
		GPUsIndex:  gpusIndex,
		GPUMemory:  totalGPUMemory,
		GPUDevices: devices,
		CPUNum:     cpuNum,
		Memory:     totalMemory,
		AvailableMemory: availableMemeory,
	}
	return
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
func GetPhysicalCPUNum() (n int32, err error) {
	physicalCnt, err := cpu.Counts(false)
	if err != nil {
		glog.Error(err)
		return
	}
	n = int32(physicalCnt)
	return
}

func (m *Machine) UpdateMemory() (err error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		glog.Error(err)
		return
	}
	m.AvailableMemory = int64(v.Available)
	return
}

func (m *Machine) UpdateCPUUtilization() (err error) {
	cpuUtilizations, err := cpu.Percent(time.Duration(0), false)
	if err != nil {
		if err != nil {
			glog.Error(err)
			return
		}
	}
	// 从百分比 转化为比例(0-1)
	m.CPUUtilization = float32(cpuUtilizations[0]) * 0.01
	return
}

func (m *Machine) UpdateGPU() (err error) {
	usedMem := int64(0)
	totalUtilization := uint(0)
	for _, dev := range m.GPUDevices {
		pInfo, err := dev.GetAllRunningProcesses()
		if err != nil {
			glog.Error(err)
			return err
		}
		for _, p := range pInfo {
			usedMem += int64(p.MemoryUsed)
		}
		status, err := dev.Status()
		if err != nil {
			glog.Error(err)
			return err
		}
		totalUtilization += *status.Utilization.GPU
	}
	m.AvailableGPUMemory = m.GPUMemory - usedMem
	if m.GPUNum > 0 {
		// 从百分比 转化为比例(0-1)
		m.CPUUtilization = float32(totalUtilization) * 0.01 / float32(m.GPUNum)
	}
	return
}

func (m *Machine) Update() (err error) {
	err = m.UpdateMemory()
	if err != nil {
		glog.Error(err)
		return
	}
	err = m.UpdateCPUUtilization()
	if err != nil {
		glog.Error(err)
		return
	}
	err = m.UpdateGPU()
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
