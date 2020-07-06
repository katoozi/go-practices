package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.URL.Path)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "the current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8080"

	r := http.NewServeMux()
	r.HandleFunc("/", myHandler)
	r.HandleFunc("/time", timeHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	// run this: go tool pprof http://localhost:1234/debug/pprof/profile
	// benchmarking: ab -k -c 10 -n 100000 http://localhost:8080/time
	// -c : concurrent requests
	// -n : number of requests
}
