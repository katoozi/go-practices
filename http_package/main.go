package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Http Package")

	// check that os.File implement io.Reader interface or not.
	var _ io.Reader = (*os.File)(nil)

	httpServer()
}

func getRequest() {
	resp, err := http.Get("https://example.com/get")
	if err != nil {
		log.Fatal("Unable to get request")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read body")
	}
	fmt.Println(string(content))
}

func uploadImagePostrequest() {
	imageData, err := os.Open("code.png")
	if err != nil {
		log.Fatal("Unable to open image file")
	}
	defer imageData.Close()

	resp, err := http.Post("https://example.com/post/image", "image/jpg", imageData)
	if err != nil {
		log.Fatal("unable to upload image")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read body")
	}
	fmt.Println(string(content))
}

func httpServer() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("/foo hanel the request")
		fmt.Fprintf(w, "Hello %s", html.UnescapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
