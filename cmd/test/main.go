package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files for the React app
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	// Start the server on port 8080
	log.Printf("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
