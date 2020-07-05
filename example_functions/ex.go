package ex

// F1 will calculate fibonacci sequence
func F1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return F1(n-1) + F1(n-2)
}

// S1 will return the s length
func S1(s string) int {
	return len(s)
}
