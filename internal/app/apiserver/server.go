package apiserver

import (
	"context"
	"encoding/json"
	"example.com/prj/model"
	"example.com/prj/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func NewServer(store store.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.ConfigureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/users", s.HandleUsersCreate()).Methods("POST")
}

func (s *Server) HandleUsersCreate() http.HandlerFunc {
	ctx := context.Background()
	type request struct {
		Email    string `json:"email"`
		Password string `json:"pwd"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{Email: req.Email, Password: req.Email}
		if err := s.store.User().Create(ctx, u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
