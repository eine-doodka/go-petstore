package test

import (
	"context"
	"example.com/prj/model"
	"example.com/prj/store"
	"github.com/go-playground/validator/v10"
)

type UserRepository struct {
	store    *Store
	users    map[string]*model.User
	validate *validator.Validate
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	if err := u.Validate(r.validate); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	r.users[u.Email] = u
	u.ID = len(r.users)
	return nil
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
