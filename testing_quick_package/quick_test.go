package main

import (
	"testing"
	"testing/quick"
)

var N = 1000000

func TestWithSystem(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == (b + a)
	}
	err := quick.Check(condition, &quick.Config{MaxCount: N})
	// quick will generate random numbers based on condition args N times.
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestWithSelf(t *testing.T) {
	condition := func(a, b uint16) bool {
		return Add(a, b) == Add(b, a)
	}
	err := quick.Check(condition, &quick.Config{MaxCount: N})
	// quick will generate random numbers based on condition args N times.
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
