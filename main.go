package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	sparkles := []Sparkle{}
	server := NewServer(&sparkles)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/sparkles.json", server.ServeHTTP)

	fmt.Printf("Listening and serving HTTP on http://%s\n", server.address)
	log.Fatal(http.ListenAndServe(server.address, nil))
}
