package main

import (
	"flag"
	"fmt"
)

func main() {
	k := flag.Bool("k", false, "k is a bool arg")
	o := flag.Int("O", 1, "O is a int arg")
	flag.Parse()

	valueK := *k
	valueO := *o

	fmt.Println(valueK)
	fmt.Println(valueO)
}
