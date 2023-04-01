package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tnaucoin/go-web-app/pkg/config"
	"github.com/tnaucoin/go-web-app/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	// Add middleware
	mux.Use(middleware.Recoverer)
	// Register route handlers
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
