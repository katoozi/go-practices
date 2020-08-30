package main

import (
	"fmt"
	"sync"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func() {
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("remove from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("adding to queue")
		queue = append(queue, i)
		go removeFromQueue()
		c.L.Unlock()
	}
}
