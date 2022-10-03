package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "%s\n", "Custom text 1")
		if err != nil {
			log.Fatal(err)
		}

		_, err = fmt.Fprintln(w, "Custom text 2")
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write([]byte("Custom text 3\n"))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
