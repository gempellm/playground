package main

import (
	"fmt"
	"log"
	"net/http"
)

type HttpHandler struct{}

func (HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "Был запрошен путь: %q\n", req.RequestURI)
}

func main() {
	const addr = "127.0.0.1:8080"
	handler := HttpHandler{}

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}
