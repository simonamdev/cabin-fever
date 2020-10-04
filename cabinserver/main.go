package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// Simplest web server example, taken from https://golang.org/doc/articles/wiki/
func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
