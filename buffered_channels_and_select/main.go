package main

import (
	"fmt"
	"time"
)

func main() {
	// buf := make(chan int, 5)
	// buf <- 3
	// buf <- 4
	// fmt.Println(len(buf))
	// buf <- 5
	// fmt.Println(len(buf))
	// fmt.Println(<-buf)
	// fmt.Println(<-buf)
	// fmt.Println(<-buf)
	// output:
	// 2
	// 3
	// 3
	// 4
	// 5
	ic := make(chan int)

	// if two channel are recive data, go will pick one of them randomly
	select {
	case v1 := <-WaitAndSend(3, 2):
		fmt.Println(v1)
	case v2 := <-WaitAndSend(5, 3):
		fmt.Println(v2)
	case ic <- 23:
		fmt.Println("ic Recived a value")
	default:
		fmt.Println("All Channels are slow")
	}
}

// WaitAndSend is a go channel generator
func WaitAndSend(v, i int) chan int {
	retChn := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		retChn <- v
	}()
	return retChn
}
