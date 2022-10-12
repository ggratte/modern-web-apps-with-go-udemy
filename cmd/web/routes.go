package main

import (
	"net/http"

	"github.com/ggratte/modern-web-apps-with-go-udemy/pkg/config"
	"github.com/ggratte/modern-web-apps-with-go-udemy/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(noSurf)
	mux.Use(sessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
