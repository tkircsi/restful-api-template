package main

import (
	"encoding/json"
	"net/http"
	"tkircsi/restful-template/pkg/models"

	"github.com/go-chi/chi"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home"))
}

func (app *application) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		app.notFound(w, r)
		return
	}
	prod, err := app.products.Get(id)
	if err != nil {
		app.notFound(w, r)
		return
	}
	jsonBytes, err := json.Marshal(prod)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (app *application) GetAll(w http.ResponseWriter, r *http.Request) {
	prods, err := app.products.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	jsonBytes, err := json.Marshal(prods)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (app *application) Save(w http.ResponseWriter, r *http.Request) {
	var prod models.Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	app.products.Save(&prod)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
