package store

import (
	"context"
	"database/sql"
)

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) Create(ctx context.Context, tx *sql.Tx, role *Role) error {
	query := `
		INSERT INTO roles (name, description, level)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)

	defer cancel()

	err := tx.QueryRowContext(
		ctx,
		query,
		role.Name,
		role.Description,
		role.Level,
	).Scan(&role.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *RoleStore) GetByName(ctx context.Context, slug string) (*Role, error) {
	query := `
		SELECT id, name, description, level 
		FROM roles 
		WHERE name = $1
	`

	role := &Role{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		slug,
	).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.Level,
	)
	if err != nil {
		return nil, err
	}

	return role, nil

}

func (s *RoleStore) Delete(ctx context.Context, tx *sql.Tx, roleID int64) error {

	query := `
		DELETE FROM roles WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, roleID)
	if err != nil {
		return err
	}

	return nil
}
