package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----------")
}

func main() {
	f, err := os.Create("traceFile.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	var mem runtime.MemStats

	printStats(mem)
	for i := 0; i < 3; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operartion Failed")
		}
	}
	printStats(mem)

	for i := 0; i < 5; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operartion Failed")
		}
		time.Sleep(time.Millisecond)
	}
	printStats(mem)

	// go tool trace traceFile.out
}
