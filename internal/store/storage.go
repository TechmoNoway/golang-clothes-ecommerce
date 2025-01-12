package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrNotFound          = errors.New("resource not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
		GetById(context.Context, int64) (*User, error)
		GetAll(context.Context) ([]User, error)
		GetByEmail(context.Context, string) (*User, error)
		Update(context.Context, *User) error
		DeleteByID(context.Context, int64) error
	}
	Roles interface {
		Create(context.Context, *Role) error
		GetByName(context.Context, string) (*Role, error)
		Delete(context.Context, *sql.Tx, int64) error
	}
	Categories interface {
		Create(context.Context, *Category) error
		GetAll(context.Context) ([]Category, error)
	}
	Products interface {
		Create(context.Context, *Product) error
		GetAll(context.Context) ([]Product, error)
		GetById(context.Context, int64) (*Product, error)
		GetAllByName(context.Context, string) ([]Product, error)
		GetAllByCategoryID(context.Context, int64) ([]Product, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Product) error
	}
	Orders interface {
		Create(context.Context, *Order) error
		GetAll(context.Context) ([]Order, error)
		GetAllByUserID(context.Context, int64) ([]Order, error)
		Update(context.Context, *Order) error
	}
	OrderItems interface {
		Create(context.Context, *OrderItem) error
		GetAll(context.Context) ([]OrderItem, error)
		GetAllByOrderID(context.Context, int64) ([]OrderItem, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:      &PostStore{db},
		Users:      &UserStore{db},
		Roles:      &RoleStore{db},
		Products:   &ProductStore{db},
		Orders:     &OrderStore{db},
		Categories: &CategoryStore{db},
		OrderItems: &OrderItemStore{db},
	}
}
