package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error while reading line: %v\n", err)
			break
		}
		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil

}

func main() {
	err := charByChar("file.txt")
	if err != nil {
		fmt.Println("Error in reading file", err)
	}
}
