package store_test

import (
	"os"
	"testing"
)

var (
	dbConnString string
)

func TestMain(m *testing.M) {
	dbConnString = os.Getenv("TEST_DATABASE_CONN_STRING")
	if dbConnString == "" {
		dbConnString = "host=localhost dbname=apiserver user=user password=user"
	}

	os.Exit(m.Run())
}
