package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	// "io"
	"net/http"

	"github.com/artemxgod/http-rest-api/internal/app/model"
	"github.com/artemxgod/http-rest-api/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

var (
	errIncorrectEmailorPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

const (
	// session name will be displayed in the respond header
	sessionName        = "KainSession"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID

)

// context key
type ctxKey int8

// server structure
type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store // sessionStore contains session for 1 month
}

// server configuration
func newServer(p_store store.Store, p_sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        p_store,
		sessionStore: p_sessionStore,
	}

	s.configureRouter()

	return s
}

// delegate s.router.ServeHTTP method to server struct
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// says which user is authenticated now
func (s *server) handleWhoami() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// we are expecting user to be logged in, so we check contect for him
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	})
}

// configuring router with handlers
func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	// allow request from any domains, by creating a header for browser
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	// only private/... URL's
	private := s.router.PathPrefix("/private").Subrouter()
	// Use() allows to apply middleware to router
	// BTW we should give cookie when request to pass MidWare
	// and given cookie have to be already in
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
}

// Creates a header with request ID
func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String() // creating random id
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

// Logging requests
func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr" : r.RemoteAddr,
			"request_id": r.Context().Value(ctxKeyRequestID),
		})
		// output when we make request
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &ResponseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		// output after request was served
		logger.Infof(
			"completed with %d %s in %v", 
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
			)
	})
}

// middle-ware component 
func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// getting cashed session by name
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		// checking if user session exists
		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		// check if user is in database
		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		// if user was find, we are going on
		// we are sending respond with contect so next time we dont have to check this user
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

// we are returning a function instead of using a function because
// before returning default function we can specify some values
func (s *server) handleUsersCreate() http.HandlerFunc {
	// specify values here
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		// io.WriteString(w, "Hello")
		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

// creates session
func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePasswowrd(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailorPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.ID
		if err = s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

// respond with error
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// respond to http request
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
