package main

import (
	"fmt"
)

func main() {
	var myInt interface{} = 120

	k, ok := myInt.(int)
	if !ok {
		fmt.Println("error in int type assertion")
	} else {
		fmt.Printf("value: %d, %T\n", k, k)
	}

	f, ok := myInt.(float64)
	if !ok {
		fmt.Println("error in float64 type assertion")
	} else {
		fmt.Printf("value: %f, %T\n", f, f)
	}

	_, ok = myInt.(bool)
	fmt.Println("parseable to bool: ", ok)
}
