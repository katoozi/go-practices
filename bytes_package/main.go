package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("some data"))
	fmt.Fprintf(&buffer, " to write")
	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)

	buffer.Reset()

	buffer.Write([]byte("this a test string!!!!"))
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if io.EOF == err {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
