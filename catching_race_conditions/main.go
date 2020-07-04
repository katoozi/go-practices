package main

import (
	"fmt"
	"sync"
)

func main() {
	var w sync.WaitGroup
	var mutex sync.Mutex
	var i int

	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < 10; i++ {
		w.Add(1)
		go func(j int) {
			defer w.Done()
			mutex.Lock() // lock the writes
			k[j] = j
			mutex.Unlock() // unlock the writes
		}(i) // pass i as argument
	}

	w.Wait()
	k[2] = 10
	fmt.Printf("k = %#v\n", k)

	// run: go run -race main.go
}
