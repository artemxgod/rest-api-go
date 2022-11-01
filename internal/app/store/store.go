package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository // let us use userrepo from outside
}

func New(p_config *Config) *Store {
	return &Store{
		config: p_config,
	}
}

// connecting to database
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	/* To check if we've connected to db we should ping it
	because actual connection is doing after first request*/
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// discconnect from database
func (s *Store) Close() {
	s.db.Close()
}

// Let us create userRepository variable
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
