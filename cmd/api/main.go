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
	http.Handle("/", server)

	fmt.Printf("Listening and serving HTTP on `%s`\n", server.Address)
	log.Fatal(http.ListenAndServe(server.Address, nil))
}
