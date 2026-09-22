package main

import (
	"fmt"
	"net/http"
)

// hello is a handler function that responds to HTTP requests with a simple "Hello, World!" message.
// It takes an http.ResponseWriter to write the response and an http.Request to access request data.
// It fulfills the http.HandlerFunc interface, allowing it to be used as a handler for HTTP requests.
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	//Registers my function as a handler for the root path "/"
	http.HandleFunc("/", hello)

	//Starts the HTTP server on port 8080 and listens for incoming requests
	http.ListenAndServe(":8080", nil)
}