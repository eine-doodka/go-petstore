package apiserver

import (
	"context"
	"example.com/prj/store/sql"
	"github.com/jackc/pgx/v4"
	"net/http"
)

func Start(config *Config) error {
	ctx := context.Background()
	db, err := NewDB(ctx, config.DbConnectionString)
	if err != nil {
		return err
	}
	defer db.Close(ctx)
	store := sql.New(db)
	srv := NewServer(store)

	return http.ListenAndServe(
		config.BindAddr,
		srv,
	)
}

func NewDB(ctx context.Context, connectionString string) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
