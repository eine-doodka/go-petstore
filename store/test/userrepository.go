package test

import (
	"context"
	"example.com/prj/model"
	"example.com/prj/store"
	"github.com/go-playground/validator/v10"
)

type UserRepository struct {
	store    *Store
	users    map[int]*model.User
	validate *validator.Validate
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	if err := u.Validate(r.validate); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	u.ID = len(r.users)
	r.users[u.ID] = u
	return nil
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) FindById(ctx context.Context, uid int) (*model.User, error) {
	u, ok := r.users[uid]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
