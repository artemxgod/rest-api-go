package teststore_test

// import (
// 	"testing"

// 	"github.com/artemxgod/http-rest-api/internal/app/model"
// 	"github.com/artemxgod/http-rest-api/internal/app/store"
// 	"github.com/artemxgod/http-rest-api/internal/app/store/teststore"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserRepository_Create(t *testing.T) {
// 	s := teststore.New()
// 	u := model.TestUser(t)

// 	err := s.User().Create(u)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, u)
// }

// func TestUserRepository_FindByEmail(t *testing.T) {
// 	s := teststore.New()
// 	email := "user@example.org"
// 	u := model.TestUser(t)
// 	u.Email = email 
// 	s.User().Create(u)
// 	u1, err := s.User().FindByEmail(email)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, u1)
// }

// func TestUserRepository_Find(t *testing.T) {
// 	s := teststore.New()
// 	u := model.TestUser(t)
// 	s.User().Create(u)

// 	u1, err := s.User().Find(u.ID)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, u1)
// }

// func TestUserRepository_FindByEmail_non(t *testing.T) {
// 	s := teststore.New()
// 	email := "user@example.org"
// 	_, err := s.User().FindByEmail(email)
// 	assert.EqualError(t, store.ErrRecordNotFound, err.Error())

// }
