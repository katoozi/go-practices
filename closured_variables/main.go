package main

import (
	"fmt"
	"time"
)

func main() {
	// Notice that closured variables in goroutines are evaluated when the goroutine actually runs and when the go statement is executed in order to create a new goroutine.
	// This means that closured variables are going to be replaced by their values when the Go scheduler decides to execute the relevant code.
	for i := 0; i <= 20; i++ {
		// i := i // this line will solve the problem.
		go func() {
			fmt.Println(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()
}
