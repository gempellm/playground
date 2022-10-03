package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/slow", func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintln(w, "ok")
	})

	server := &http.Server{Addr: "127.0.0.1:8080"}

	go func() {
		srvErr := server.ListenAndServe() // Стартуем сервер и блокируем горутину
		if srvErr != http.ErrServerClosed {
			log.Fatalf("server.ListenAndServe(): %v", srvErr)
		}
		fmt.Println("Server successfully stopped")
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	termSig := <-termChan // Блокируем работу основной горутины тут

	log.Println("Gracefully shutdown starter with signal:", termSig)

	closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(closeCtx)
	if err != nil {
		log.Println("server.Shutdown() failed:", err)
	}

	fmt.Println("Server successfully shutdown")
}
