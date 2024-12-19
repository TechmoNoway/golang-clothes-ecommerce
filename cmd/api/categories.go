package main

import (
	"net/http"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

func (app *application) getAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
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

type CreateCategoryPayload struct {
	CategoryName string `json:"category_name" validate:"required,max=100"`
}

func (app *application) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateCategoryPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	category := &store.Category{
		CategoryName: payload.CategoryName,
	}

	ctx := r.Context()

	err = app.store.Categories.Create(ctx, category)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusCreated, category)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
