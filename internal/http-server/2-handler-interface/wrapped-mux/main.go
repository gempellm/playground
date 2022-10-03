package main

import "net/http"

type httpHandler struct {
	mux *http.ServeMux
}

func newHandler() *httpHandler {
	handler := &httpHandler{
		mux: http.NewServeMux(),
	}

	handler.mux.HandleFunc("/", handler.index)
	return handler
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (httpHandler) index(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("index\n"))
}

func main() {
	mux := newHandler()
	http.ListenAndServe(":8080", mux)
}
