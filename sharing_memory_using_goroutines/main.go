package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value := newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go monitor()

	var w sync.WaitGroup

	for r := 0; r < 100; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * 100))
		}()
	}
	w.Wait()
	fmt.Println("Last value: ", read())
}
