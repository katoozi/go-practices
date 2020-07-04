package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	signal := make(chan bool)
	go func() {
		for i = 0; i <= 255; i++ {
			// this loop will never stop
		}
		signal <- true
	}()
	fmt.Println("Leaving goroutine!")
	runtime.Gosched()
	runtime.GC()
	<-signal

	fmt.Println("Good byte!")
}
