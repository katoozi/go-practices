package main

import (
	"fmt"
	"sort"
)

func main() {
	var firstArray = [5]int{5, 2, 4, 5, 6}

	firstSlice := make([]int, 2, 4)

	twoDemintionalSlice := make([][]int, 4)
	fmt.Println("two demintional slice ", twoDemintionalSlice)

	fmt.Printf("firstArray type is %T\n", firstArray)
	fmt.Println(firstArray)
	fmt.Printf("firstArray type is %T\n", firstSlice)
	fmt.Println(firstSlice)
	fmt.Println(cap(firstSlice))
	fmt.Println(len(firstSlice))

	fmt.Println("Next Section")

	anArray := [4]int{-1, -2, -3, -4}
	refArray := anArray[:]

	fmt.Println("anArray", anArray)
	fmt.Println("refArray", refArray)
	refArray[2] = -100 // this one will change the base array. because slice is a refrence to the main array.
	fmt.Println("anArray", anArray)
	fmt.Println("refArray", refArray)

	// sorting
	unsortdSlice := []int{5, 2, 4, 5, 6}
	fmt.Println("unsortdSlice", unsortdSlice)
	sort.Slice(unsortdSlice[:], func(i, j int) bool {
		return unsortdSlice[i] < unsortdSlice[j]
	})
	fmt.Println("unsortdSlice", unsortdSlice)

	//appending
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}

	ref := a[:]
	fmt.Println("ref: ", ref)
	t := append(s, ref...)
	fmt.Println("s+ref: ", t)
	s = append(s, ref...)
	fmt.Println("s+ref: ", s)
	s = append(s, s...)
	fmt.Println("s+s: ", s)
}
