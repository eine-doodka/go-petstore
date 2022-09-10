package sql

import (
	"context"
	"example.com/prj/model"
	"example.com/prj/store"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	store    *Store
	validate *validator.Validate
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	if err := u.Validate(r.validate); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.
		store.
		dbConn.
		QueryRow(ctx,
			"INSERT INTO users(email, crypto) VALUES ($1, $2) RETURNING id",
			u.Email,
			u.EncryptedPassword).
		Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.dbConn.QueryRow(ctx,
		"SELECT id, email, crypto FROM users WHERE email = $1",
		email).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
