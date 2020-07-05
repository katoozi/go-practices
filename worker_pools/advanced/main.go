package main

import (
	"fmt"
	"sync"
	"time"
)

// Client is simulating incomming clients
type Client struct {
	id, integer int
}

// Data is the type we used in channel
type Data struct {
	job    Client
	sqaure int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		sqaure := c.integer * c.integer
		output := Data{c, sqaure}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity os clients:", cap(clients))
	fmt.Println("Capacity os data:", cap(data))

	nJobs := 50
	nWorkers := 5

	go create(nJobs)

	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("CLient ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsqaure: %d\n", d.job.integer, d.sqaure)
		}
		finished <- true
	}()
	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
