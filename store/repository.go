package store

import (
	"context"
	"example.com/prj/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindById(ctx context.Context, uid int) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
