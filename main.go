package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go Server!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running at http://localhost:8111")
	http.ListenAndServe(":8111", nil)
}
