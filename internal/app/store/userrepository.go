package store

// UserRepository is responible for operations with table "users"

import (
	"github.com/artemxgod/http-rest-api/internal/app/model"
)

// Contains link to our main store
type UserRepository struct {
	store *Store
}

/* Creates new record in database, returns modified User struct that contains ID
Requires user struct with email and password*/
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	пш
	if err := r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
	u.Email, u.EncryptedPassowrd,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

/* Let us find the record by email, returns record or error
Requires email address*/
func (r * UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, encrypted_password FROM users WHERE email = $1",
	 email).Scan(
		&u.ID, &u.Email, &u.EncryptedPassowrd); err != nil {
			return nil, err
		}
	return u, nil
}