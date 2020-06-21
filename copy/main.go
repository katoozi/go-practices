package main

import "fmt"

func main() {
	a1 := [9]int{1, 34, 234, -323, -3132, -12, -234234, 8323123, 123456789}
	a2 := []int{-1, -34, -234, 323, 3132, 12, 234234, -8323123, -123456789}

	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)

	copy(a2, a1[0:]) // a1 converted to slice

	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
}
