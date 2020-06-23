package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

type XYZ struct {
	X int
	Y int
	Z int
}

type myStructure struct {
	Name string
}

func newMyStrcuture(name string) *myStructure {
	return &myStructure{name}
}

func main() {
	var a1 aStructure
	a1.height = 100
	a1.weight = 200
	a1.person = "user"
	fmt.Println(a1)

	p1 := aStructure{person: "p_user", height: 200, weight: 100}
	fmt.Println(p1.person)

	var s1 XYZ = XYZ{X: 100}
	fmt.Println(s1.X, s1.Y, s1.Z)

	pSlice := [4]XYZ{}
	pSlice[0] = s1
	fmt.Println(pSlice)

	// pointers to structures
	my := newMyStrcuture("pointer_user")
	fmt.Println(*my)
}
