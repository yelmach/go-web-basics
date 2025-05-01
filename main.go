package main

import (
	"fmt"
	"net/http"
	handle "web/Handler"
)

func main() {
	// route handlers
	http.HandleFunc("/", handle.IndexHandler)
	http.HandleFunc("/ascii-art", handle.AsciiArtHandler)
	http.HandleFunc("/export", handle.AsciiArtHandler)


	// Serve static files
	http.Handle("/Static/", http.FileServer(http.Dir("./")))
	
	// Start the server
	fmt.Println("The server is running on http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}
