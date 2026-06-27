package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About page")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/about", about)

	fmt.Println("Server running at http://localhost:8111")
	http.ListenAndServe(":8111", nil)
}
