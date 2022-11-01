package apiserver

import (
	"io"
	"net/http"

	"github.com/artemxgod/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// Initialize http server with config logger and mux
func New(p_config *Config) *APIServer {
	return &APIServer{
		config: p_config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Starting our initialized http server
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Configure logger with certain level of logging
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// Configure out router handlers
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
} 

// we are returning a function instead of using a function because
// before returning default function we can specify some values
func (s *APIServer) handleHello() http.HandlerFunc {
	// specify values here
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
