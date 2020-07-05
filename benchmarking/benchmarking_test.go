package main

import "testing"

// this technique is used to prevent the compiler from performing any optimizations that will exclude
// the function that you want to measure from being executed because its results are never used.
var result int

func benchmarkFibo1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo1(n)
	}
	result = r
}

func benchmarkFibo2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo2(n)
	}
	result = r
}

func benchmarkFibo3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibo3(n)
	}
	result = r
}

func Benchmark30Fibo1(b *testing.B) {
	benchmarkFibo1(b, 30)
}
func Benchmark30Fibo2(b *testing.B) {
	benchmarkFibo2(b, 30)
}
func Benchmark30Fibo3(b *testing.B) {
	benchmarkFibo3(b, 30)
}

func Benchmark50Fibo1(b *testing.B) {
	benchmarkFibo1(b, 50)
}
func Benchmark50Fibo2(b *testing.B) {
	benchmarkFibo2(b, 50)
}
func Benchmark50Fibo3(b *testing.B) {
	benchmarkFibo3(b, 50)
}

// run benchmarks: go test -bench=. ./...
// include memory allocation statistics in output: go test -benchmem -bench=. ./...
