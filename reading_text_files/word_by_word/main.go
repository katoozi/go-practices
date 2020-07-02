package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func wordByWord(file string) error {
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
		for index, word := range strings.Fields(line) {
			fmt.Println(index, word)
		}
		fmt.Println()
	}
	return nil

}

func main() {
	err := wordByWord("file.txt")
	if err != nil {
		fmt.Println("Error in reading file", err)
	}
}
