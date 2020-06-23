package main

import "fmt"

func main() {
	// A rune is an int32 value, and therefore a Go type, that is used for representing a Unicode code point.
	// A Unicode code point, or code position, is a numerical value that is usually used for representing single Unicode characters;
	// however, it can also have alternative meanings, such as providing formatting information.
	// NOTE: You can consider a string as a collection of runes.
	const r1 = '$'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %d\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)

	aStaring := []byte("test")
	for x, y := range aStaring {
		fmt.Println(x, y)
		fmt.Printf("Char: %c\n", aStaring[x])
	}
}
