package sqlstore

import (
	"database/sql"

	"github.com/artemxgod/http-rest-api/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	// config         *Config
	db             *sql.DB
	userRepository *UserRepository // let us use userrepo from outside
}

func New(p_db *sql.DB) *Store {
	return &Store{
		db: p_db,
	}
}



// Let us create userRepository variable
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
