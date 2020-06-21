package main

import "fmt"

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	// There are two main reasons for using pointers in your programs:
	// 1. Pointers allow you to share data, especially between Go functions.
	// 2. Pointers can be extremely useful when you want to differentiate between a zero value and a value that is not set.
	i := -10
	j := 100

	pI := &i
	pJ := &j

	fmt.Println("PI memory: ", pI)
	fmt.Println("PJ memory: ", pJ)
	fmt.Println("Pi value: ", *pI)
	fmt.Println("PJ value: ", *pJ)

	*pI = 123456
	*pI--
	fmt.Println("i: ", i)

	getPointer(pJ)
	fmt.Println("j: ", j)

	k := returnPointer(12)
	fmt.Println("k memory: ", k)
	fmt.Println("k value: ", *k)
}
