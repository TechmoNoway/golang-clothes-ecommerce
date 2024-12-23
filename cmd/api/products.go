package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

type productKey string

const productCtx productKey = "product"

func (app *application) getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := app.store.Products.GetAll(r.Context())
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

func (app *application) getAllProductsByNameHandler(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("name")
	products, err := app.store.Products.GetAllByName(r.Context(), productName)
	fmt.Println(products)
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

func (app *application) getAllProductsByCategoryIDHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.ParseInt(r.URL.Query().Get("categoryID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	products, err := app.store.Products.GetAllByCategoryID(r.Context(), categoryID)
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

type CreateProductPayload struct {
	ProductName string `json:"product_name" validate:"required,max=100"`
	Description string `json:"content" validate:"required,max=1000"`
	Price       int64  `json:"price" validate:"required"`
	Stock       int64  `json:"stock"`
	Size        string `json:"size"`
	Color       string `json:"color"`
	CategoryID  int64  `json:"category_id"`
}

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateProductPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	product := &store.Product{
		ProductName: payload.ProductName,
		Description: payload.Description,
		Price:       payload.Price,
		Stock:       payload.Stock,
		Size:        payload.Size,
		Color:       payload.Color,
		CategoryID:  payload.CategoryID,
	}

	ctx := r.Context()

	err = app.store.Products.Create(ctx, product)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusCreated, product)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdateProductPayload struct {
	ProductName *string `json:"product_name" validate:"omitempty,max=100"`
	Description *string `json:"description" validate:"omitempty,max=1000"`
	Price       *int64  `json:"price"`
	Stock       *int64  `json:"stock"`
	Size        *string `json:"size"`
	Color       *string `json:"color"`
	CategoryID  *int64  `json:"category_id"`
}

func (app *application) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	product := app.getProductFromCtx(r)

	var payload UpdateProductPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.ProductName != nil {
		product.ProductName = *payload.ProductName
	}

	if payload.Description != nil {
		product.Description = *payload.Description
	}

	if payload.Price != nil {
		product.Price = *payload.Price
	}

	if payload.Stock != nil {
		product.Stock = *payload.Stock
	}

	if payload.Size != nil {
		product.Size = *payload.Size
	}

	if payload.Color != nil {
		product.Color = *payload.Color
	}

	if payload.CategoryID != nil {
		product.CategoryID = 1
	}

	ctx := r.Context()

	err = app.updateProduct(ctx, product)
	if err != nil {
		app.internalServerError(w, r, err)
	}

	err = app.jsonResponse(w, http.StatusOK, product)
	if err != nil {
		app.internalServerError(w, r, err)
	}

}

func (app *application) deleteProductHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getProductFromCtx(r *http.Request) *store.Product {
	product, _ := r.Context().Value(productCtx).(*store.Product)
	return product
}

func (app *application) updateProduct(ctx context.Context, product *store.Product) error {
	err := app.store.Products.Update(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
