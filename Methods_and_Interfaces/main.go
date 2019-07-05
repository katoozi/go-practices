package main

import "fmt"

// GlobalInterface is just a simple interface
type GlobalInterface interface {
	Say()
	Increment()
	GetInternalValue() int
}

type Data struct {
	i int
}

func (d *Data) Say() {
	fmt.Println("Hello")
}

func (d *Data) Increment() {
	d.i++
}

func (d *Data) GetInternalValue() int {
	return d.i
}

// CallGetInternalValue is a function that call interface GetInternalValue method
func CallGetInternalValue(obj GlobalInterface) int {
	return obj.GetInternalValue()
}

func main() {
	data := &Data{}
	data2 := &Data{}
	data.Increment()
	data2.Increment()
	data2.Increment()
	fmt.Println(CallGetInternalValue(data))
	fmt.Println(CallGetInternalValue(data2))
}
