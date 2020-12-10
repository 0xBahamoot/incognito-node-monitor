package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type systemInfo struct {
	MemUsed uint64
	// Disk     int
	// DiskUsed int
	CPU float64
}

var system systemInfo

func monitorSystem() {
	for {
		time.Sleep(2 * time.Second)
		memory, err := memory.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		// fmt.Printf("memory total: %d bytes\n", memory.Total)
		// fmt.Printf("memory used: %d bytes\n", memory.Used)
		system.MemUsed = memory.Used
		before, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		total := float64(after.Total - before.Total)
		// fmt.Printf("cpu: %f %%\n", float64(after.User-before.User+after.System-before.System)/total*100)
		system.CPU = float64(after.User-before.User+after.System-before.System) / total * 100
	}

}
