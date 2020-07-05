package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// fibonacci recursive one
func fibo1(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	}
	time.Sleep(time.Millisecond)
	return int64(fibo1(n-1)) + int64(fibo1(n-2))
}

// fibonacci with memorization technique
func fibo2(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	time.Sleep(50 * time.Millisecond)
	return fn[n]
}

// N1 will find out  the given integer is a prime number or not.
func N1(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

// N2 will find out  the given integer is a prime number or not.
func N2(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	cpuFile, err := os.Create("cpuProfile.out")
	if err != nil {
		fmt.Println("Error while creating cpuProfile.out:", err)
		return
	}
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	total := 0
	for i := 2; i < 100000; i++ {
		n := N1(i)
		if n {
			total++
		}
	}
	fmt.Println("Total Primes:", total)

	total = 0
	for i := 2; i < 100000; i++ {
		n := N2(i)
		if n {
			total++
		}
	}
	fmt.Println("Total Primes:", total)

	for i := 1; i < 10; i++ {
		n := fibo1(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	for i := 1; i < 90; i++ {
		n := fibo2(i)
		fmt.Print(n, " ")
	}
	fmt.Println()

	runtime.GC()

	memory, err := os.Create("memoryProfile.out")
	if err != nil {
		fmt.Println("Error while creating memoryProfile.out:", err)
		return
	}
	defer memory.Close()

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed.")
		}
		time.Sleep(50 * time.Millisecond)
	}
	err = pprof.WriteHeapProfile(memory)
	if err != nil {
		fmt.Println(err)
		return
	}

	// first run: go run main.go
	// cpu profiling: go tool pprof cpuProfile.out
	// memory profiling: go tool pprof memoryProfile.out
	// web view: go tool pprof -http=localhost:8080 cpuProfile.out
}
