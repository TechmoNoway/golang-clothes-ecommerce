package main

import (
	"net/http"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

func (app *application) getAllCategoriesHanler(w http.ResponseWriter, r *http.Request) {
	categories, err := app.store.Categories.GetAll(r.Context())
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	err = app.jsonResponse(w, http.StatusOK, categories)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}
