package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer) // recover from panic
	router.Use(app.enableCORS)

	router.Get("/", app.Home)

	router.Post("/authenticate", app.Authenticate)

	router.Get("/movie/{id:[0-9]+}", app.Movie)

	router.Get("/movies", app.AllMovies)

	router.Post("/add", app.AddMovie)

	router.Delete("/delete/movie/{id:[0-9]+}", app.DeleteMovie)

	return router
}
