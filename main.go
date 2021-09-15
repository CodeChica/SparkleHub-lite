package main

import (
	"log"
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
	log.Fatal(server.Run(listenAddress()))
}
