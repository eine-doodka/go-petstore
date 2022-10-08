package test

import (
	"example.com/prj/model"
	"example.com/prj/store"
	"github.com/go-playground/validator/v10"
)

type Store struct {
	userRepo *UserRepository
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}
	s.userRepo = &UserRepository{
		store:    s,
		users:    make(map[int]*model.User),
		validate: validator.New(),
	}
	return s.userRepo
}
