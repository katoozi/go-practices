package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "some example string"

	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bs) // print the hexadecimal
}
