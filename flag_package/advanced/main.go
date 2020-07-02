package main

import (
	"flag"
	"fmt"
	"strings"
)

// NamesFlag is a structure
type NamesFlag struct {
	Names []string
}

// GetNames will return the names slice
func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

// Set wil fill the NamesFlag struct
func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}
	names := strings.Split(v, ",")
	for _, name := range names {
		s.Names = append(s.Names, name)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	k := flag.Bool("k", true, "k is bool arg")
	o := flag.Int("o", 1, "o is int arg")
	flag.Var(&manyNames, "names", "Comma-separated list of names")

	flag.Parse()

	fmt.Println("-k ", *k)
	fmt.Println("-o ", *o)

	fmt.Println("Names: ")
	for i, name := range manyNames.GetNames() {
		fmt.Println("\t", i, ":", name)
	}

	fmt.Println("Remaining command line arguments:")
	for index, value := range flag.Args() {
		fmt.Println(index, value)
	}
}
