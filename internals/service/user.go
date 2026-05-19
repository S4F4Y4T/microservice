package service

import (
	"microservice/internals/model"
	"net/http"
)

type UserService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	// Implement logic to write users to the response
}
