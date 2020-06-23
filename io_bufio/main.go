package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("google.com")

	io.WriteString(os.Stdout, "download\n")

	mystring := ""
	args := os.Args
	if len(args) == 1 {
		mystring = "Please Send one Argument."
	} else {
		mystring = args[1]
	}
	io.WriteString(os.Stdout, mystring)

	var f *os.File

	f = os.Stdin

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}
