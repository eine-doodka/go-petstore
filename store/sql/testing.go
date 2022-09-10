package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
	"testing"
)

func TestDB(t *testing.T, connStr string) (*pgx.Conn, func(...string)) {
	t.Helper()
	ctx := context.Background()
	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(ctx); err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(ctx, fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		db.Close(ctx)
	}
}
