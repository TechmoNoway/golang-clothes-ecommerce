package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	QueryTimeoutDuration = time.Second * 100
	ErrNotFound          = errors.New("resource not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		GetById(context.Context, int64) (*User, error)
		GetAll(context.Context) ([]User, error)
	}
	Roles interface {
		Create(context.Context, *sql.Tx, *Role) error
		GetByName(context.Context, string) (*Role, error)
		Delete(context.Context, *sql.Tx, int64) error
	}
	Products interface {
		Create(context.Context, *sql.Tx, *Product) error
		GetAll(context.Context) ([]Product, error)
		GetById(context.Context, int64) (*Product, error)
		GetAllByName(context.Context, string) ([]Product, error)
		Delete(context.Context, int64) error
	}
	Orders interface {
		Create(context.Context, *sql.Tx, *Order) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Roles:    &RoleStore{db},
		Products: &ProductStore{db},
		Orders:   &OrderStore{db},
	}
}
