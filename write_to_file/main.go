package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	s := []byte("some data to write!")

	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Error in write f1", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Error in open f2", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	if err != nil {
		fmt.Println("Error in write f2", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println("Error in open f3", err)
		return
	}
	defer f3.Close()
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	if err != nil {
		fmt.Println("Error in write f3", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println("Error is write f4", err)
		return
	}

	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println("Error in opening f5", err)
		return
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println("Error in write f5", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
