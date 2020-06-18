package main

import "fmt"

func main() {
	var firstArray = [5]int{5, 2, 4, 5, 6}

	firstSlice := make([]int, 2, 4)

	fmt.Printf("firstArray type is %T\n", firstArray)
	fmt.Println(firstArray)
	fmt.Printf("firstArray type is %T\n", firstSlice)
	fmt.Println(firstSlice)
	fmt.Println(cap(firstSlice))
	fmt.Println(len(firstSlice))
}
