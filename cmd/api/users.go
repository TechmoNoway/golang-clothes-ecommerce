package main

import (
	"net/http"
	"strconv"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	user, err := app.getUser(r.Context(), userID)
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

	err = app.jsonResponse(w, http.StatusOK, user)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := app.store.Users.GetAll(r.Context())
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

	err = app.jsonResponse(w, http.StatusOK, users)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
