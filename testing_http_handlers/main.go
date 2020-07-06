package main

import (
	"fmt"
	"net/http"
)

func checkStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `Fine!`)
}

func statusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Serving: %s\n", r.Host)
}

func main() {
	PORT := ":8080"

	http.HandleFunc("/check-status-ok", checkStatusOK)
	http.HandleFunc("/check-status-not-found", statusNotFound)
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
