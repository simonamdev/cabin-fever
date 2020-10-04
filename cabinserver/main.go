package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The time now is %s", time.Now().String())
	log.Println(message)
	// Allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, message)
}

// Simplest web server example, taken from https://golang.org/doc/articles/wiki/
func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
