package service

import (
	"encoding/json"
	"fmt"
	"gopet/storage"
	"net/http"
	"strconv"
	"sync"
)

type Service struct {
	userStorage *storage.Storage
	mu          sync.RWMutex
}

func New(storage *storage.Storage) *Service {
	return &Service{
		userStorage: storage,
	}
}

func (s *Service) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi there")
}

func (s *Service) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user storage.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age <= 0 {
		http.Error(w, "Name is required and age must be positive", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.userStorage.CreateUser(user)
	defer s.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	s.mu.RLock()

	user, ok := s.userStorage.GetUser(id)
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding user data", http.StatusInternalServerError)
		return
	}

	defer s.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //return the user data as JSON
	w.Write(j)
}

func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	s.mu.Lock()

	if ok := s.userStorage.DeleteUser(id); !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	defer s.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
