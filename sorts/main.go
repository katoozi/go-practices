package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	ints := []int{5, 6, 2, 4, 6, 1, 9, 4}

	sort.Strings(strs)
	fmt.Println("Sorted: ", strs)
	fmt.Println("Is Sorted: ", sort.StringsAreSorted(strs))

	sort.Ints(ints)
	fmt.Println("Sorted: ", ints)
	fmt.Println("Is Sorted: ", sort.IntsAreSorted(ints))
}
