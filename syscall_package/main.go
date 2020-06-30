package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid is : ", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("user id: ", uid)

	message := []byte{'H', 'E', 'L', 'L', 'O', '!', '\n'}
	fd := 1                    // file descriptor
	syscall.Write(fd, message) // fmt package print functions are implemented like this

	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
