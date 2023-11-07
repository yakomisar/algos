package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"time"
)

func main() {
	cpuFile, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	memFile, err := os.Create("mem.pprof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memFile.Close()
	// Простая функция для имитации нагрузки на CPU и выделения памяти.
	for i := 0; i < 1000000; i++ {
		s := make([]byte, 1000)
		if i%1000 == 0 {
			time.Sleep(500 * time.Millisecond)
		}
		for i := 0; i < 1000; i++ {
			s[i] = byte(i % 256)
		}
	}

	pprof.WriteHeapProfile(memFile)
}
