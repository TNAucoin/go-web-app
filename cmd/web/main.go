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
	app.UseCache = false
	// assign template cache to the app config
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// pass the app config to the renderer
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
