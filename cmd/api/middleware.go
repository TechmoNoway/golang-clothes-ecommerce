package main

import (
	"context"
	"fmt"

	"github.com/TechmoNoway/golang-clothes-ecommerce/internal/store"
)

// User handlers
func (app *application) getUserById(ctx context.Context, userID int64) (*store.User, error) {
	user, err := app.store.Users.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (app *application) getAllUsers(ctx context.Context) ([]store.User, error) {
	users, err := app.store.Users.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(users)

	return users, nil
}

func (app *application) updateUser(ctx context.Context) error {
	return nil
}

func (app *application) deleteUser(ctx context.Context) error {
	return nil
}

// Products
func (app *application) getAllProducts(ctx context.Context) error {
	return nil
}

func (app *application) getProductsByName(ctx context.Context) ([]store.Product, error) {
	return nil, nil
}

func (app *application) updateProduct(ctx context.Context) error {
	return nil
}

func (app *application) deleteProduct(ctx context.Context) error {
	return nil
}
