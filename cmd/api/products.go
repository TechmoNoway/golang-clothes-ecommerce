package main

import (
	"net/http"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

func (app *application) getAllProductsHanler(w http.ResponseWriter, r *http.Request) {
	products, err := app.getAllProducts(r.Context())
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

	err = app.jsonResponse(w, http.StatusOK, products)
	if err != nil {
		app.internalServerError(w, r, err)
	}

}

func (app *application) createProductHanler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateProductHanler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteProductHanler(w http.ResponseWriter, r *http.Request) {

}
