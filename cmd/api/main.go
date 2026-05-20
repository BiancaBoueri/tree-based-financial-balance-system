package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}
