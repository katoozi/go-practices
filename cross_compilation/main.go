package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler)
	fmt.Println(" on a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())

	// complie the application
	// GOARCH: amd64, arm, 386
	// GOOS: linux, windows, darwin(for macos)
	// linux: GOOS=linux GOARCH=386 go build -o main.o main.go
	// macos: GOOS=darwin GOARCH=amd64 go build -o main.o main.go
	// windows: GOOS=windows GOARCH=amd64 go build -o app.exe main.go
}
