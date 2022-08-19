package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, connStr string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.ConnStr = connStr
	s := New(config)
	ctx := context.Background()
	if err := s.Open(ctx); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			_, err := s.dbConn.Exec(ctx, fmt.Sprintf("TRUNCATE %s CASCADE",
				strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal(err)
			}
		}
		s.Close(ctx)
	}
}
