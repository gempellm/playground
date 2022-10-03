package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,

		Timeout: 2 * time.Millisecond,
	}

	res, err := httpClient.Get("https://example.com/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Statuscode:", res.StatusCode)
}
