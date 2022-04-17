package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codechica/SparkleHub-lite/pkg/web"
)

var (
	address string
	help    bool
)

func init() {
	flag.BoolVar(&help, "help", false, "")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	flag.StringVar(&address, "address", ":"+port, "the address to bind to")
	flag.Parse()
}

func main() {
	if help == true {
		flag.Usage()
		os.Exit(0)
	} else {
		fmt.Printf("Listening and serving HTTP on `%s`\n", address)
		log.Fatal(http.ListenAndServe(address, web.NewServer(nil)))
	}
}
