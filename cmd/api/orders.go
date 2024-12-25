package main

import (
	"net/http"
	"strconv"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

type orderKey string

const orderCtx orderKey = "order"

type CreateOrderPayload struct {
	UserID     int64  `json:"user_id"`
	TotalPrice int64  `json:"total_price"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateOrderPayload
	err := readJSON(w, r, payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(payload)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	order := &store.Order{
		UserID:     payload.UserID,
		TotalPrice: payload.TotalPrice,
		Status:     payload.Status,
	}

	ctx := r.Context()

	err = app.store.Orders.Create(ctx, order)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.jsonResponse(w, http.StatusCreated, order.ID)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := app.store.Orders.GetAll(r.Context())
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

	err = app.jsonResponse(w, http.StatusOK, orders)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getAllOrdersByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	orders, err := app.store.Orders.GetAllByUserID(r.Context(), userID)
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

	err = app.jsonResponse(w, http.StatusOK, orders)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

type UpdateOrderPayload struct {
	TotalPrice *int64  `json:"total_price"`
	Status     *string `json:"status"`
}

func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	order := app.getOrderFromCtx(r)

	var payload UpdateOrderPayload
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

	if payload.TotalPrice != nil {
		order.TotalPrice = *payload.TotalPrice
	}

	if payload.Status != nil {
		order.Status = *payload.Status
	}

	ctx := r.Context()

	err = app.store.Orders.Update(ctx, order)
	if err != nil {
		app.internalServerError(w, r, err)
	}

	err = app.jsonResponse(w, http.StatusOK, order)
	if err != nil {
		app.internalServerError(w, r, err)
	}

}

func (app *application) deleteOrderHandler(w http.ResponseWriter, r *http.Request){
	
}

func (app *application) getOrderFromCtx(r *http.Request) *store.Order {
	order, _ := r.Context().Value(orderCtx).(*store.Order)
	return order
}
