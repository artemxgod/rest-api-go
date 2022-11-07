package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/artemxgod/http-rest-api/internal/app/store/sqlstore"
)

// starting the server
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	s := newServer(store)

	return http.ListenAndServe(config.BindAddr, s)
}

//initializing database
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	// ping to check connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}