// +build 386

//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2020] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package runtime

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/lastbackend/lastbackend/internal/pkg/models"
	"github.com/lastbackend/lastbackend/internal/util/system"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const MinContainerMemory = 32

func (r Runtime) NodeInfo() models.NodeInfo {

	var (
		info = models.NodeInfo{}
	)

	osInfo := system.GetOsInfo()
	hostname, err := os.Hostname()
	if err != nil {
		_ = fmt.Errorf("get hostname err: %s", err)
	}

	ip, err := system.GetHostIP(envs.Get().GetConfig().Network.Interface)
	if err != nil {
		_ = fmt.Errorf("get ip err: %s", err)
	}

	info.Hostname = hostname
	info.ExternalIP = ip
	info.OSType = osInfo.GoOS
	info.OSName = fmt.Sprintf("%s %s", osInfo.OS, osInfo.Core)
	info.Architecture = osInfo.Platform

	net := envs.Get().GetNet()
	if net != nil {
		nt := net.Info(context.Background())
		info.InternalIP = nt.IP
		info.CIDR = nt.CIDR
	}

	return info
}

func (r Runtime) NodeStatus() models.NodeStatus {

	var state = models.NodeStatus{}

	state.Capacity = NodeCapacity()
	state.Allocated = NodeAllocation()

	return state
}

func (r Runtime) NodeCapacity() models.NodeResources {

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		_ = fmt.Errorf("get memory err: %s", err)
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		_ = fmt.Errorf("get cpu err: %s", err)
	}

	var storage int64

	wd, err := os.Getwd()
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	_, _, err = c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(wd))), uintptr(unsafe.Pointer(&storage)))

	m := vmStat.Total

	return models.NodeResources{
		Storage:    int64(storage),
		CPU:        int64(cpuStat[0].Mhz) * int64(1e6) * int64(cpuStat[0].Cores),
		RAM:        int64(m),
		Pods:       int(m / MinContainerMemory),
		Containers: int(m / MinContainerMemory),
	}
}

func (r Runtime) NodeAllocation() models.NodeResources {

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		_ = fmt.Errorf("get memory err: %s", err)
	}

	m := vmStat.Free
	s := envs.Get().GetState().Pods()

	return models.NodeResources{
		RAM:        int64(m),
		CPU:        0, // TODO: need get cpu resource value
		Pods:       s.GetPodsCount(),
		Containers: s.GetContainersCount(),
	}
}
