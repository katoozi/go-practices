package main

import (
	"fmt"
	"net"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	// Every regular expression is compiled into a recognizer by building a generalized transition diagram called a finite automaton.
	// A finite automaton can be either deterministic or nondeterministic.
	// Nondeterministic means that more than one transition out of a state can be possible for the same input.
	// A recognizer is a program that takes a string x as input and is able to tell whether x is a sentence of a given language.
	// A grammar is a set of production rules for strings in a formal language.
	// The production rules describe how to create strings from the alphabet of the language that are valid according to the syntax of the language.
	// A grammar does not describe the meaning of a string or what can be done with it in whatever context – it only describes its form.
	// What is important here is to realize that grammars are at the heart of regular expressions because without a grammar,
	//  you cannot define or use a regular expression.
	ip := findIP("this 1.1.1.1 ip is connecting")
	trial := net.ParseIP(ip)
	fmt.Println(trial.To4())
}
