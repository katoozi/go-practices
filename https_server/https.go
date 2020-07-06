package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is an https server.")
}

func main() {
	http.HandleFunc("/", defaultHandler)

	err := http.ListenAndServeTLS(":1443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
