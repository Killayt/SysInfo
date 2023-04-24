package main

import (
	"fmt"
	"time"

	memory "systemMonitor/nvidia_memory"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {

	for {
		// retrieve CPU load as a percentage
		cpuLoad, err := cpu.Percent(0, false)
		if err != nil {
			fmt.Println("Error retrieving CPU load:", err)
			continue
		}

		// retrieve memory usage as a percentage
		memStat, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error retrieving memory usage:", err)
			continue
		}
		memUsage := memStat.UsedPercent

		// retrieve GPU load as a percentage (Nvidia only)
		gpuLoad := memory.GetGPU()

		// print the metrics
		fmt.Printf("CPU: %.2f%%\n", cpuLoad[0])
		fmt.Printf("RAM: %.2f%%\n", memUsage)
		fmt.Print("GPU: ", gpuLoad+"%")

		fmt.Println("\n---------------")

		// wait for 1 second before checking again
		time.Sleep(time.Second)
	}
}
