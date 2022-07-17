package main

import (
	"fmt"
	"net/http"
)

func main() {
	NewServer()
}

func NewServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}
