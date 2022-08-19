package store_test

import (
	"context"
	"example.com/prj/model"
	"example.com/prj/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, dbConnString)
	defer teardown("users")
	ctx := context.Background()
	u, err := s.User().Create(ctx, &model.User{
		Email: "user@example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, dbConnString)
	defer teardown("users")
	ctx := context.Background()
	email := "user@example.org"
	_, err := s.User().FindByEmail(ctx, email)
	assert.Error(t, err)

	s.User().Create(ctx, &model.User{
		Email:             email,
		EncryptedPassword: "CRYPTO",
	})
	res, err := s.User().FindByEmail(ctx, email)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
