package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("download from google by download manager")
	lr := io.LimitReader(r, 4) // lr will have the first 4 characters from r Reader.

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatalf("Error Happend: %s", err)
	}
}
