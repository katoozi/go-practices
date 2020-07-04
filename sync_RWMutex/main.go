package main

import (
	"fmt"
	"sync"
	"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var password = secret{password: "myPassword"}

func change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	defer c.RWM.RUnlock()

	fmt.Println("show")
	time.Sleep(3 * time.Second)
	return c.password
}

func showWithLock(c *secret) string {
	c.RWM.Lock()
	defer c.RWM.Unlock()

	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	return c.password
}

func main() {
	// The sync.RWMutex type is another kind of mutex – actually,
	// it is an improved version of sync.Mutex
	// Now, let us talk about how sync.RWMutex improves sync.Mutex.
	// Although only one function is allowed to perform write operations using a sync.RWMutex mutex,
	// you can have multiple readers owning a sync.RWMutex mutex.
	// However, there is one thing that you should be aware of: until all of the readers of a sync.RWMutex mutex unlock that mutex,
	// you cannot lock it for writing, which is the small price you have to pay for allowing multiple readers.
	var showFunction = show

	var waitGroup sync.WaitGroup

	// sync.RWMutex mutex is much faster than the version that uses sync.Mutex.

	fmt.Println("Pass: ", showFunction(&password))

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass: ", showFunction(&password))
		}()
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		change(&password, "1233423423")
	}()

	waitGroup.Wait()
	fmt.Println("Pass: ", showFunction(&password))
}
