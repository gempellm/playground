package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			_, err := fmt.Fprintln(w, "Hello, World!")
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}

		testHandler := func(w http.ResponseWriter, req *http.Request) {
			_, err := fmt.Fprintln(w, "This is a /test!")
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}

		http.HandleFunc("/", helloHandler)
		http.HandleFunc("/test", testHandler)

		err := http.ListenAndServe(":8080", nil)
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
