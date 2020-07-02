package main

import "fmt"

type twoInts struct {
	X, Y int64
}

func (a twoInts) method(b twoInts) twoInts {
	// a is a receiver
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func main() {
	two := twoInts{10, 0}
	fmt.Println(two)
}
