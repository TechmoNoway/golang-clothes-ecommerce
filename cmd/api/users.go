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

type UpdateUserPayload struct {
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	AvatarUrl *string `json:"avatar_url"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	Address   *string `json:"address"`
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.getUserFromContext(r)

	var payload UpdateUserPayload
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

	if payload.Username != nil {
		user.Username = *payload.Username
	}

	if payload.Email != nil {
		user.Email = *payload.Email
	}

	if payload.AvatarUrl != nil {
		user.AvatarUrl = *payload.AvatarUrl
	}

	if payload.FirstName != nil {
		user.FirstName = *payload.FirstName
	}

	if payload.LastName != nil {
		user.LastName = *payload.LastName
	}

	if payload.Phone != nil {
		user.Phone = *payload.Phone
	}

	if payload.Address != nil {
		user.Address = *payload.Address
	}

	ctx := r.Context()

	err = app.store.Users.Update(ctx, user)
	if err != nil {
		app.internalServerError(w, r, err)
	}

	err = app.jsonResponse(w, http.StatusCreated, user)
	if err != nil {
		app.internalServerError(w, r, err)
	}

}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
