package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserPayload struct {
	Username  string `json:"username" validate:"required,max=100"`
	Email     string `json:"email" validate:"required,email,max=255"`
	Password  string `json:"password" validate:"required,min=3,max=72"`
	FirstName string `json:"first_name" validate:"required,min=3,max=72"`
	LastName  string `json:"last_name" validate:"required,min=3,max=72"`
	Phone     string `json:"phone" validate:"required,min=3,max=10"`
	Address   string `json:"address" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
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

	user := &store.User{
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  payload.Password,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Phone:     payload.Phone,
		Address:   payload.Address,
		RoleID:    1,
	}

	ctx := r.Context()

	plainToken := uuid.New().String()

	//TODO: this is for email activation
	// hash := sha256.Sum256([]byte(plainToken))
	// hashToken := hex.EncodeToString(hash[:])

	err = app.store.Users.Create(ctx, user)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestResponse(w, r, err)
		case store.ErrDuplicateUsername:
			app.badRequestResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	userWithToken := UserWithToken{
		User:  user,
		Token: plainToken,
	}

	err = app.jsonResponse(w, http.StatusCreated, userWithToken)
	if err != nil {
		fmt.Println(err)
		app.internalServerError(w, r, err)
	}

}

type LoginUserTokenPayload struct {
	Username string `json:"username" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserTokenPayload
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

	user, err := app.store.Users.GetByUsername(r.Context(), payload.Username)
	fmt.Println("Get username error")
	fmt.Println(err)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.unauthorizedErrorResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	fmt.Println(user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	fmt.Println("Get compare password error")

	fmt.Println(err)
	if err != nil {
		app.unauthorizedBasicErrorResponse(w, r, err)
		return
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(app.config.auth.token.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.token.iss,
		"aud": app.config.auth.token.iss,
	}

	token, err := app.authenticator.GenerateToken(claims)
	fmt.Println("Generate token err")
	fmt.Println(err)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	fmt.Print(token)

	err = app.jsonResponse(w, http.StatusCreated, token)
	fmt.Println(err)
	if err != nil {
		app.internalServerError(w, r, err)
	}
}
