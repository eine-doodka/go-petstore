package apiserver

import (
	"context"
	"example.com/prj/store/sql"
	sessions2 "github.com/gorilla/sessions"
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
	sessions := sessions2.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessions)

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
