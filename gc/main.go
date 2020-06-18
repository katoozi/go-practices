package main

import (
	"fmt"
	"runtime"
)

func printStates(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("-------------")
}

func main() {
	fmt.Println("Garbage Collector")
	var mem runtime.MemStats
	fmt.Println(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation Failed")
		}
		printStates(mem)
	}
	printStates(mem)
}
