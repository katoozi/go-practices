package main

import (
	"fmt"
	"os"
	s "strings"
	"unicode"
)

func main() {
	var f = fmt.Printf
	f("to upper %s\n", s.ToUpper("Hello world"))
	f("to lower %s\n", s.ToLower("Hello world"))

	f("%s\n", s.Title("hello world"))

	f("%v\n", s.EqualFold("hello World", "HELLO WORld"))
	f("%v\n", s.EqualFold("hello World", "HELLo worl"))

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))

	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))

	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Split("abcd efg", " "))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
	f("%s\n", s.Replace("a b c d e f g", " ", "_", -1))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))

	str := s.NewReader("this is an error!!")
	fmt.Println("str length: ", str.Len())
	n, err := str.WriteTo(os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nwrote %d bytes to Stdout\n", n)
}
