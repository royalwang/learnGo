package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:           "localhost:8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	handler := &myHandler{}
	http.Handle("/foo", handler)
	server.ListenAndServe()
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
