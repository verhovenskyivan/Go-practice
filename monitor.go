package main

import (
    "fmt"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "time"
)

func main() {
    // Monitor every 1 second
    interval := time.Second

    for {
        // Get CPU usage
        cpuPercent, _ := cpu.Percent(interval, false)
        fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])

        // Get RAM usage
        memInfo, _ := mem.VirtualMemory()
        fmt.Printf("RAM Usage: Used %.2f GB, Total %.2f GB, %.2f%% Used\n",
            float64(memInfo.Used)/1024/1024/1024,
            float64(memInfo.Total)/1024/1024/1024,
            memInfo.UsedPercent)

        time.Sleep(interval)
    }
}
