package main

import (
	"fmt"
	"log"
	"net/http"

	"mokhan.ca/CodeChica/sparkleapi/pkg/domain"
)

func main() {
	sparkles := []domain.Sparkle{}
	server := NewServer(&sparkles)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/sparkles.json", server.ServeHTTP)

	fmt.Printf("Listening and serving HTTP on `%s`\n", server.address)
	log.Fatal(http.ListenAndServe(server.address, nil))
}
