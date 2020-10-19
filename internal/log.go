package internal

import (
	"fmt"
	"runtime"
)

func Log(s string, args ...interface{}) {
	fmt.Printf("  => "+s+"\n", args...)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("\t---\n")
	fmt.Printf("\tAlloc = %v MiB\t", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\t", bToMb(m.TotalAlloc))
	fmt.Printf("\tStackInuse = %v MiB\t", bToMb(m.StackInuse))
	fmt.Printf("\tStackSys = %v MiB\t", bToMb(m.StackSys))
	fmt.Printf("\tSys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("\tHeapInuse = %v MiB\t", bToMb(m.HeapInuse))
	fmt.Printf("\tHeapSys = %v MiB\t", bToMb(m.HeapSys))
	fmt.Printf("\tHeapIdle = %v MiB\t", bToMb(m.HeapIdle))
	fmt.Printf("\tHeapReleased = %v MiB\t", bToMb(m.HeapReleased))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Printf("\t---\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}