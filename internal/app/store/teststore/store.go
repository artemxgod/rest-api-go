package teststore

import (
	"github.com/artemxgod/http-rest-api/internal/app/model"
	"github.com/artemxgod/http-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository // let us use userrepo from outside
}

func New() *Store {
	return &Store{}
}


// Let us create userRepository variable
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
