package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux = &http.ServeMux{}
	mux = new(http.ServeMux)

	mux.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "One")
	})

	mux.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Two")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
