package main

import "fmt"

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Printf("x: %x\n", sLiteral)
	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q: %+q\n", sLiteral)
	fmt.Printf(" x: % x\n", sLiteral)
	fmt.Printf("s: %s\n", sLiteral)

	s2 := "€£3"
	fmt.Println(len(sLiteral))
	fmt.Println(s2)
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)
	fmt.Printf("s3 length: %d\n", len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
