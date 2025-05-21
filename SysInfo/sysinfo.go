package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	fmt.Println("=== System Information ===")
	hostname, err := host.Info()
	if err != nil {
		fmt.Printf("Error fetching hostname: %v\n", err)
		return
	}
	fmt.Printf("Hostname: %s\n", hostname.Hostname)
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	if hostname != nil {
		fmt.Printf("OS Version: %s\n", hostname.PlatformVersion)
	}
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error fetching CPU info: %v\n", err)
		return
	}
	fmt.Printf("\nCPU Information:\n")
	if len(cpuInfo) > 0 {
		fmt.Printf("  Model: %s\n", cpuInfo[0].ModelName)
		fmt.Printf("  Cores: %d\n", cpuInfo[0].Cores)
	}
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Error fetching CPU usage: %v\n", err)
		return
	}
	if len(percent) > 0 {
		fmt.Printf("  Usage: %.2f%%\n", percent[0])
	}
	vm, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error fetching memory info: %v\n", err)
		return
	}
	fmt.Printf("\nMemory Information:\n")
	fmt.Printf("  Total: %.2f GB\n", float64(vm.Total)/1024/1024/1024)
	fmt.Printf("  Used: %.2f GB\n", float64(vm.Used)/1024/1024/1024)
	fmt.Printf("  Free: %.2f GB\n", float64(vm.Free)/1024/1024/1024)
	fmt.Printf("  Usage: %.2f%%\n", vm.UsedPercent)

	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Printf("Error fetching disk partitions: %v\n", err)
		return
	}
	fmt.Printf("\nDisk Information:\n")
	for _, partition := range partitions {
		if partition.Mountpoint == "C:\\" {
			usage, err := disk.Usage(partition.Mountpoint)
			if err != nil {
				fmt.Printf("Error fetching disk usage for %s: %v\n", partition.Mountpoint, err)
				continue
			}
			fmt.Printf("  Drive: %s\n", partition.Mountpoint)
			fmt.Printf("  Total: %.2f GB\n", float64(usage.Total)/1024/1024/1024)
			fmt.Printf("  Used: %.2f GB\n", float64(usage.Used)/1024/1024/1024)
			fmt.Printf("  Free: %.2f GB\n", float64(usage.Free)/1024/1024/1024)
			fmt.Printf("  Usage: %.2f%%\n", usage.UsedPercent)
			break
		}
	}
}