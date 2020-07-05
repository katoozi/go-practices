package main

import "fmt"

// Add is function
func Add(x, y uint16) uint16 {
	var i uint16
	for i = 0; i < x; i++ {
		y++
	}
	return y
}

func main() {
	fmt.Println(Add(0, 0))
}
