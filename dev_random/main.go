package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	if err != nil {
		fmt.Println("error while opening /dev/random", err)
		return
	}

	var seed int64
	err = binary.Read(f, binary.LittleEndian, &seed)
	if err != nil {
		fmt.Println("Error while reading binary", err)
		return
	}
	fmt.Println(seed)
}
