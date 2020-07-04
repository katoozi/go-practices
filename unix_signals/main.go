package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("HandleSignal() Caught: ", signal)
}

func main() {
	// The SIGKILL and SIGSTOP signals cannot be caught, blocked, or ignored.
	// SIGKILL: 9
	// signal.SIGINFO in not available on Linux machines
	// The most common way to send a signal to a process is by using the kill(1) utility.
	// By default, kill(1) sends the SIGTERM signal.

	// NOTE: handle all signals but respond only to the ones that really interest you.
	// This is a much better and safer technique.
	sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, os.Interrupt, syscall.SIGINFO)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGINFO:
				handleSignal(sig)
				return
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(1)
			}
		}
	}()
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
