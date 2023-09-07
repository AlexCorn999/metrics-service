package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	router *chi.Mux
}

func NewAPIServer() *APIServer {
	return &APIServer{
		router: chi.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureRouter(); err != nil {
		return err
	}
	return http.ListenAndServe(":8080", s.router)
}

func (s *APIServer) ConfigureRouter() error {
	return nil
}
