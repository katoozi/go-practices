package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Printf("Receive Job %d\n", j)
			} else {
				fmt.Println("Receive All jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("Send a job", j)
	}
	close(jobs)

	fmt.Println("Send All Jobs")

	<-done

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Printf("Queue %s\n", elem)
	}

}
