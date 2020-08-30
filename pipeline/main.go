package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multiplyStream <- i * multiplier:
				}
			}
		}()
		return multiplyStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, addictive int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addStream <- i + addictive:
				}
			}
		}()
		return addStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2) // i*2+1*2

	for v := range pipeline {
		fmt.Println(v)
	}
}
