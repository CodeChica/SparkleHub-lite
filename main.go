package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func listenAddress() string {
	return ":" + port()
}

func main() {
	sparkles := []Sparkle{}

	server := NewServer(&sparkles)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/sparkles.json", server.ServeHTTP)

	address := listenAddress()
	fmt.Printf("Listening and serving HTTP on http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
