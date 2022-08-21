package store

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	config   *Config
	dbConn   *pgx.Conn
	userRepo *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, s.config.ConnStr)
	if err != nil {
		return err
	}
	if err := conn.Ping(ctx); err != nil {
		return err
	}
	s.dbConn = conn
	return nil
}

//func (s *Store) Migrate(ctx context.Context) error {
//	s.dbConn.
//}

func (s *Store) Close(ctx context.Context) {
	s.dbConn.Close(ctx)
}

func (s *Store) User() *UserRepository {
	if s.userRepo == nil {
		s.userRepo = &UserRepository{
			store:    s,
			validate: validator.New(),
		}
	}
	return s.userRepo
}
