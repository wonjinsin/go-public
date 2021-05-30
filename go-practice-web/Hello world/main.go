package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Bar!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world")
	})

	http.HandleFunc("/bar", barHandler)

	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":5000", nil)
}
