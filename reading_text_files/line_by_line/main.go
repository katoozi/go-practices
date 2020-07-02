package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		// line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error while reading line: %v\n", err)
			break
		}
		// fmt.Println(line)
		fmt.Println(string(line))
	}
	return nil

}

func main() {
	err := lineByLine("file.txt")
	if err != nil {
		fmt.Println("Error in reading file", err)
	}
}
