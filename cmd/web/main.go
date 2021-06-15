package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/config"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/handlers"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCatche()
	if err != nil {
		log.Fatal("Can't create Template catche")
	}
	app.TemplateCatche = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Aplication open on port %s", port)

	_ = http.ListenAndServe(port, nil)
}
