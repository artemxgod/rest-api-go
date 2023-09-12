package sqlstore_test

// import (
// 	"testing"

// 	"github.com/artemxgod/http-rest-api/internal/app/model"
// 	"github.com/artemxgod/http-rest-api/internal/app/store"
// 	"github.com/artemxgod/http-rest-api/internal/app/store/sqlstore"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserRepository_Create(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")
// 	s := sqlstore.New(db)
// 	u := model.TestUser(t) 

// 	err := s.User().Create(u)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, u)
// }

// func TestUserRepository_FindByEmail(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")
// 	s := sqlstore.New(db)

// 	s.User().Create(model.TestUser(t))


// 	email := "user@example.org"
// 	u, err := s.User().FindByEmail(email)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, u)
// }

// func TestUserRepository_Find(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")

// 	s := sqlstore.New(db)
// 	u1 := model.TestUser(t)
// 	s.User().Create(u1)


// 	u, err := s.User().Find(u1.ID)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, u)
// }

// func TestUserRepository_FindByEmail_non(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("users")
// 	s := sqlstore.New(db)

// 	email := "user@example.org"
// 	_, err := s.User().FindByEmail(email)
// 	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

// }
