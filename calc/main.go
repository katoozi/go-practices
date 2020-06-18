package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give me the max and min numbers")
		os.Exit(1)
	}
	arguments := os.Args

	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[2], 64)

	fmt.Println(min, max)

	log.Println("Download")
}

// Fib will caculate the fibunachi
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
