package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gen(min, max int, createNUmber chan<- int, end chan bool) {
	for {
		select {
		case createNUmber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	createNumber := make(chan int)
	end := make(chan bool)

	go gen(0, 112831723123, createNumber, end)

	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", <-createNumber)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Existing...")
	end <- true
}
