package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type getQuoteContract struct {
	Quote string `json:"quote"`
}

func main() {
	ctx := context.Background()
	contract, err := getQuote(ctx)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Kanye West says: %s\n", contract.Quote)
}

func getQuote(ctx context.Context) (*getQuoteContract, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.kanye.rest/", nil)
	httpClient := &http.Client{Timeout: 5 * time.Second}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contract := new(getQuoteContract)
	err = json.Unmarshal(data, contract)
	if err != nil {
		return nil, err
	}

	return contract, nil
}
