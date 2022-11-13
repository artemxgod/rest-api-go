package store

// store package  provides us only with interfaces 
// real structs and methods are in the sqlstore and sql test

import "github.com/artemxgod/http-rest-api/internal/app/model"

// User repo interface
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

