package server

import (
	"gopet/service"
	"net/http"
)

type Server struct {
	service *service.Service
}

func New(svc *service.Service) *Server {
	return &Server{service: svc}
}

func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", s.service.HelloHandler)
	mux.HandleFunc("POST /users", s.service.CreateUser)
	mux.HandleFunc("GET /users/{id}", s.service.GetUser)
	mux.HandleFunc("DELETE /users/{id}", s.service.DeleteUser)

	return http.ListenAndServe(addr, mux)
}
