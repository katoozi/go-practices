package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch Time: ", time.Now().Unix(), time.Now().UnixNano())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time Diffrence: ", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)

	loc, _ := time.LoadLocation("Asia/Tehran")
	tehranTime := t.In(loc)
	fmt.Println("Tehran: ", tehranTime)
}
