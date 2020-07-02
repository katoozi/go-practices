package main

import (
	"fmt"
	"math"
)

// Shape is a interface for geometric chapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

type sqaure struct {
	X float64
}

func (s sqaure) Area() float64 {
	return s.X * s.X
}
func (s sqaure) Perimeter() float64 {
	return 4 * s.X
}

type circle struct {
	R float64
}

func (s circle) Area() float64 {
	return s.R * s.R * math.Pi
}
func (s circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

// Calculate will calculate the Area and Perimeter
func Calculate(v Shape) {
	// _, ok := v.(circle)
	// if ok {
	// 	fmt.Println("Is a circle")
	// }
	// _, ok = v.(sqaure)
	// if ok {
	// 	fmt.Println("Is a sqaure")
	// }
	fmt.Println(v.Area())
	fmt.Println(v.Perimeter())
}

// tellInterface will calculate the Area and Perimeter
func tellInterface(v interface{}) {
	switch objType := v.(type) {
	case circle:
		fmt.Println("Is a circle")
	case sqaure:
		fmt.Println("Is a sqaure")
	default:
		fmt.Printf("this object is a %T\n", objType)
	}
}

func main() {
	// Putting it simply, interfaces should be utilized when there is a need for making sure
	// that certain conditions will be met and certain behaviors will be anticipated from a Go element.
	x := sqaure{10}
	y := circle{20}

	tellInterface(x)
	Calculate(x)
	tellInterface(y)
	Calculate(y)

	tellInterface(10)
	tellInterface(10.00)
	tellInterface(complex(1, 2))
	tellInterface(complex(40, 3))
	tellInterface('f')
	var char int = 163
	tellInterface(char)
	tellInterface("golang")
	tellInterface([]byte("asdasd"))
}
