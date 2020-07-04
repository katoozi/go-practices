package main

import (
	"fmt"
	"time"
)

var times int

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum += i
		}
		c <- sum
	case <-f:
		return
	}
}

func main() {
	cc := make(chan chan int)

	for i := 0; i < 10; i++ {

		// this code also works without f because goroutine doesn't have any for loop
		// The type of a signal channel can be anything you want, including bool, which is used in the preceding code, and struct{},
		// The main advantage of a struct{} signal channel is that no data can be sent to it,
		// which can save you from bugs and misconceptions.
		f := make(chan bool)

		go f1(cc, f)
		ch := <-cc
		ch <- i
		for sum := range ch {
			fmt.Printf("Sum(%d)=%d", i, sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}
