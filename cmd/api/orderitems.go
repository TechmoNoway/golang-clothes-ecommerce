package main

import (
	"net/http"
	"strconv"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

type orderItemKey string

const orderItemCtx orderItemKey = "orderItem"

func (app *application) getAllOrderItemsHandler(w http.ResponseWriter, r *http.Request) {

	orderItems, err := app.store.OrderItems.GetAll(r.Context())
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

	err = app.jsonResponse(w, http.StatusOK, orderItems)
	if err != nil {
		app.internalServerError(w, r, err)
	}

}

type CreateOrderItemPayload struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
	Price     int64 `json:"price"`
}

func (app *application) getAllOrderItemsByOrderIDHandler(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.ParseInt(r.URL.Query().Get("orderID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	orderItems, err := app.store.OrderItems.GetAllByOrderID(r.Context(), orderID)
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

	err = app.jsonResponse(w, http.StatusOK, orderItems)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) createOrderItemHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateOrderItemPayload
	err := readJSON(w, r, &payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	orderItem := &store.OrderItem{
		OrderID:   payload.OrderID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		Price:     payload.Price,
	}

	ctx := r.Context()

	err = app.store.OrderItems.Create(ctx, orderItem)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusCreated, orderItem)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
