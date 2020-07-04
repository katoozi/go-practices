package main

import (
	"context"
	"fmt"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	// The main purpose of the context package is to define the Context type and support cancellation.
	// f1(4)
	// f2(4)
	// f3(4)

	// f1(2)
	// f2(2)
	// f3(2)

	// f1(10)
	// f2(10)
	// f3(10)
}
