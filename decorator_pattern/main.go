package main

import (
	"log"
)

type Object func(int) int

func LogDecorator(fn Object) Object {
	return func(a int) int {
		log.Println("Start")
		result := fn(a)
		log.Println("End")
		return result
	}
}

func Double(n int) int {
	return n * 2
}

func main() {
	f := LogDecorator(Double)
	log.Println(f(5))
}
