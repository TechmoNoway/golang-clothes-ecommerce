package main

import (
	"context"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	user, err := app.store.Users.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
