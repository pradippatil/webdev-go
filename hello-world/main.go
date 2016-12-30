package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server running on port 8080. Open http://localhost:8080/World in browser")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
