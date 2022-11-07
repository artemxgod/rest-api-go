package apiserver

import (
	"io"
	"net/http"

	"github.com/artemxgod/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// server structure
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

// server configuration
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}
	
	s.configureRouter()

	return s
}

// delegate s.router.ServeHTTP method to server struct
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configuring router with handlers
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.HandleUsersCreate()).Methods("POST")
}

// we are returning a function instead of using a function because
// before returning default function we can specify some values
func (s *server) HandleUsersCreate() http.HandlerFunc {
	// specify values here
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}