package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://example.com/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status code: %v\n", res.StatusCode)
}
