package main

import (
	"fmt"
	"os"
)

func main() {
	info, _ := os.Stat("file.txt")
	fmt.Println(info.Mode().String()[1:10])
}
