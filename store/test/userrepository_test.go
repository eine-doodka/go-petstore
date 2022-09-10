package test_test

import (
	"context"
	"example.com/prj/model"
	"example.com/prj/store"
	"example.com/prj/store/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := test.NewStore()
	ctx := context.Background()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(ctx, u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := test.NewStore()
	ctx := context.Background()

	email := "user@example.org"
	_, err := s.User().FindByEmail(ctx, email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	err = s.User().Create(ctx, u)
	assert.NoError(t, err)

	u, err = s.User().FindByEmail(ctx, email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
