package main

func main() {
	// Generally speaking, constants are global variables, so you might rethink your approach
	// if you find yourself defining too many constant variables with a local scope.
	// NOTE: Strictly speaking, the value of a constant variable is defined at compile time not at run time.
	const HEIGHT = 200
	const (
		C1 = "c1"
		C2 = "c2"
		C3 = "c3"
	)
	// Go compiler considers the results of all operations applied to constants as constants.
}
