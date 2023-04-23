package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	// // check if Nvidia GPU is available on the system
	// nvidiaGPU, err := nvidia.NewGpu()
	// hasNvidia := err
	// if err != nil {
	// 	errors.New("You system have no video-adapters")
	// }

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

		// // retrieve GPU load as a percentage (Nvidia only)
		// var gpuLoad float64
		// if hasNvidia {
		// 	gpuStat, err := nvidia.UsedMemory()
		// 	if err != nil {
		// 		fmt.Println("Error retrieving GPU load:", err)
		// 		continue
		// 	}
		// 	gpuLoad = float64(gpuStat) / 100
		// }

		// print the metrics
		fmt.Printf("CPU: %.2f%%\n", cpuLoad[0])
		fmt.Printf("RAM: %.2f%%\n", memUsage)
		// if hasNvidia {
		// 	fmt.Printf("GPU: %.2f%%\n", gpuLoad)
		// }

		fmt.Println("---------------")

		// wait for 1 second before checking again
		time.Sleep(time.Second)
	}
}
