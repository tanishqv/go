package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tanishqv/go/mywebapp-go/pkg/config"
	"github.com/tanishqv/go/mywebapp-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Creating HTTP handler, often called a "mux" or "multiplexer"
	mux := chi.NewRouter()

	// Installing middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
