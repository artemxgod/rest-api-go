package store_test

import (
	"testing"

	"github.com/artemxgod/http-rest-api/internal/app/model"
	"github.com/artemxgod/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")
	s.User().Create(model.TestUser(t))

	email := "user@example.org"
	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail_non(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
	
}