package sql

import (
	"example.com/prj/store"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	dbConn   *pgx.Conn
	userRepo *UserRepository
}

func New(db *pgx.Conn) *Store {
	return &Store{
		dbConn: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepo == nil {
		s.userRepo = &UserRepository{
			store:    s,
			validate: validator.New(),
		}
	}
	return s.userRepo
}
