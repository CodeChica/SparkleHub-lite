package main

import (
	"fmt"
	"log"
	"net/http"

	"mokhan.ca/CodeChica/sparkleapi/pkg/domain"
	"mokhan.ca/CodeChica/sparkleapi/pkg/web"
)

func main() {
	sparkles := []domain.Sparkle{}
	server := web.NewServer(&sparkles)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/sparkles.json", server.ServeHTTP)

	fmt.Printf("Listening and serving HTTP on `%s`\n", server.Address)
	log.Fatal(http.ListenAndServe(server.Address, nil))
}
