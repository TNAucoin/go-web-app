package main

import (
	"fmt"
	"github.com/tnaucoin/go-web-app/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting App on port:%s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
