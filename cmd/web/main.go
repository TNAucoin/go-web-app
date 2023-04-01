package main

import (
	"fmt"
	"github.com/tnaucoin/go-web-app/pkg/config"
	"github.com/tnaucoin/go-web-app/pkg/handlers"
	"github.com/tnaucoin/go-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// Setup App config
	app.UseCache = false
	app.TemplateCache = tc
	// Create handlers repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// Setup Render
	render.NewTemplates(&app)

	fmt.Printf("Starting App on port:%s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
