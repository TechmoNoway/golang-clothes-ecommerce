package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	AvatarUrl string `json:"avatar_url"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	RoleID    int64  `json:"role_id"`
	Role      Role   `json:"role"`
	CreatedAt string `json:"created_at"`
}

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password, first_name, last_name, phone, address, role_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at
	`

	// ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	// defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var userRole int
	role := user.Role.Name
	if role == "user" || role == "" {
		userRole = 1
	}

	err = s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		hash,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Address,
		userRole,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		fmt.Println("db error")
		return err
	}

	return nil
}

func (s *UserStore) GetById(ctx context.Context, userID int64) (*User, error) {
	query := `
		SELECT users.id, username, email, password, avatar_url, first_name, last_name, phone, address, created_at, roles.*
		FROM users
		JOIN roles ON (users.role.id = roles.id) 
		WHERE users.id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.AvatarUrl,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Address,
		&user.CreatedAt,
		&user.Role.ID,
		&user.Role.Name,
		&user.Role.Description,
		&user.Role.Level,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStore) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `
		SELECT users.id, username, email, password, avatar_url, first_name, last_name, phone, address, created_at, roles.*
		FROM users
		JOIN roles ON (users.role.id = roles.id) 
		WHERE users.username = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.AvatarUrl,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Address,
		&user.CreatedAt,
		&user.Role.ID,
		&user.Role.Name,
		&user.Role.Description,
		&user.Role.Level,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStore) GetAll(ctx context.Context) ([]User, error) {

	query := `
		SELECT users.id, username, email, password, avatar_url, first_name, last_name, phone, address, created_at, roles.*
		FROM users
		JOIN roles ON (users.role_id = roles.id) 
		WHERE 1 = 1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)

	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userList []User
	for rows.Next() {
		var newUser User
		err := rows.Scan(
			&newUser.ID,
			&newUser.Username,
			&newUser.Email,
			&newUser.Password,
			&newUser.AvatarUrl,
			&newUser.FirstName,
			&newUser.LastName,
			&newUser.Phone,
			&newUser.Address,
			&newUser.CreatedAt,
			&newUser.Role.ID,
			&newUser.Role.Name,
			&newUser.Role.Description,
			&newUser.Role.Level,
		)
		if err != nil {
			return nil, err
		}

		userList = append(userList, newUser)
	}

	return userList, nil
}

func (s *UserStore) DeleteByID(ctx context.Context, userID int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Update(ctx context.Context, user *User) error {
	query := `
	UPDATE users SET username, email, password, avatar_url, first_name, last_name, phone, address 
	SET username = $1, email = $2, avatar_url = $3, first_name = $4, last_name = $5, phone = $5, address = $6
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)

	defer cancel()

	_, err := s.db.ExecContext(ctx, query, user.Username, user.Email, user.AvatarUrl, user.FirstName, user.LastName, user.Phone, user.Address)
	if err != nil {
		return err
	}

	return nil
}
