package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	myURL string = "https://httpstat.us/408"
	delay int    = 1
	w     sync.WaitGroup
)

type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer w.Done()
	data := make(chan myData, 1)

	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", myURL, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		}
		pack := myData{response, err}
		data <- pack
	}()
	select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("the request was canceled")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select: ", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server response: %s\n", realHTTPData)
	}
	return nil
}

func main() {
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	w.Add(1)
	go connect(c)
	w.Wait()
	fmt.Println("Existing...")
}
