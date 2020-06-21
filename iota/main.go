package main

import "fmt"

// DIGIT is exactly the integer type.
type DIGIT int

// POWER2 is exactly the integer type.
type POWER2 int

func main() {
	const (
		ZERO DIGIT = iota
		ONE
		TWO
		THREE
		FOUR
	)
	fmt.Println(ZERO)
	fmt.Println(ONE)
	fmt.Println(TWO)
	fmt.Println(THREE)
	fmt.Println(FOUR)

	const (
		p2_0 POWER2 = 1 << iota
		_           // _ skip unwanted values
		p2_2
		_
		p2_4
		_
		p2_6
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
