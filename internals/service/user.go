package service

import (
	"context"
	"fmt"
	"microservice/internals/model"
	"net/http"
)

type UserService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(c context.Context) ([]model.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}

	return users, nil
}

func (s *UserService) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Production implementation
	// user, err := s.repo.GetUserByID()
	// if err != nil {
	// 	http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(user); err != nil {
	// 	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	// 	return
	// }
}
