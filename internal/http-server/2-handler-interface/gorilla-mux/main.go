package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	router *mux.Router
}

func newHandler() *httpHandler {
	router := mux.NewRouter()
	handler := &httpHandler{router: router}
	router.HandleFunc("/index", handler.index).Methods(http.MethodGet)
	return handler
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (httpHandler) index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {
	handler := newHandler()
	http.ListenAndServe(":8080", handler)
}
