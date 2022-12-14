package teststore

import (

	"github.com/artemxgod/http-rest-api/internal/app/model"
	"github.com/artemxgod/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

/* Creates new record in database, returns modified User struct that contains ID
Requires user struct with email and password*/
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	
	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// Search for record by ID
func (r *UserRepository) Find(ID int) (*model.User, error) {
	u, ok := r.users[ID]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// Search for record by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
		
	return nil, store.ErrRecordNotFound
}