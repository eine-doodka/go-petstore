package store

import (
	"context"
	"example.com/prj/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := r.
		store.
		dbConn.
		QueryRow(ctx,
			"INSERT INTO users(email, crypto) VALUES ($1, $2) RETURNING id",
			u.Email,
			u.EncryptedPassword).
		Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
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
		return nil, err
	}
	return u, nil
}
