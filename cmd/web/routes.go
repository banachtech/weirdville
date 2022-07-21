package main

import (
	"net/http"

	"banachtech.github.com/weirdville/cmd/pkg/config"
	"banachtech.github.com/weirdville/cmd/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter() // create router

	mux.Use(middleware.Recoverer) // gracefully handle panic errors
	mux.Use(NoSurf)               // custom middleware
	mux.Use(LoadSession)          // custom middleware
	// attach handlers
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
