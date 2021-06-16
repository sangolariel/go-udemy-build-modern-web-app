package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/config"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/handlers"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/render"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	//Env
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCatche()
	if err != nil {
		log.Fatal("Can't create Template catche")
	}
	app.TemplateCatche = tc
	app.UseCatche = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	fmt.Printf("Aplication open on port %s", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
