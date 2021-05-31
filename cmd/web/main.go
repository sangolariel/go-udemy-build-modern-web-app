package main

import (
	"fmt"
	"net/http"

	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Aplication open on port %s", port)

	_ = http.ListenAndServe(port, nil)
}