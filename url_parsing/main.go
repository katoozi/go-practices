package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

func main() {
	uri := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Error while parsing url: %v\n", err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Host)
	fmt.Println(u.Fragment)
	fmt.Println(u.Path)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
