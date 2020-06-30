package main

import "fmt"

func oneByOne(message string, s ...int) {
	fmt.Printf("Message %s\n", message)
	for index, value := range s {
		fmt.Println(index, " -> ", value)
	}
}

func main() {
	// ...int => pack data
	// data... => unpack data array|slice
	data := []int{-1, -2, -3, -4, -5, -6, 21323, 32, 213}
	oneByOne("pack numbers", 1, 2, 3, 4, 5, 6, -21323, -32, -213)
	oneByOne("unpack data variable", data...)
}
