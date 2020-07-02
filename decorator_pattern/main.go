package main

import (
	"log"
)

// Object is a func
type Object func(int) int

// LogDecorator is a decorator
func LogDecorator(fn Object) Object {
	return func(a int) int {
		log.Println("Start")
		result := fn(a)
		log.Println("End")
		return result
	}
}

// Double will calculate double of an integer
func Double(n int) int {
	return n * 2
}

func main() {
	f := LogDecorator(Double)
	log.Println(f(5))
}
