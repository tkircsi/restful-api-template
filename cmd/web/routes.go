package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(setHeaders)

	r.Get("/", app.Home)
	r.Get("/products", app.GetAll)
	r.Get("/products/{id}", app.Get)
	r.Post("/products", app.Save)
	r.NotFound(http.HandlerFunc(app.notFound))

	return r
}
