package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

var port = ":1443"

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	caCert, err := ioutil.ReadFile("client.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	srv := &http.Server{
		Addr:      port,
		Handler:   &handler{},
		TLSConfig: cfg,
	}

	fmt.Println("Listening to port number:", port)
	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))

}
