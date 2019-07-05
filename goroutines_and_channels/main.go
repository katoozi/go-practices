package main

import (
	"fmt"
	"time"
)

func main() {
	// rmemebr in this type of comunication both go routine will wait for send and recive from channel
	// quitSignal := make(chan bool)
	// go func() {
	// 	fmt.Println("Go Routines 1")
	// 	quitSignal <- true
	// }()
	// fmt.Println("main routine")
	// returnValue := <-quitSignal
	// fmt.Println(returnValue)

	ic := make(chan int)
	go periodSend(ic)
	for i := range ic {
		fmt.Println(i)
	}

	// check channel is closed or not.
	if _, ok := <-ic; ok == false {
		fmt.Println("Channel is closed")
	}

}

func periodSend(ic chan int) {
	i := 0
	for i <= 5 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic)
}
