package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tnaucoin/go-web-app/pkg/config"
	"github.com/tnaucoin/go-web-app/pkg/handlers"
	"github.com/tnaucoin/go-web-app/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// Global app config
var app config.AppConfig
var session *scs.SessionManager

func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// Setup App config
	app.UseCache = false
	app.TemplateCache = tc
	// Set Production mode
	app.InProduction = false

	// Create a new session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// Store the session in app config
	app.Session = session

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
