package service

import (
	"context"
	"fmt"
	"microservice/internals/model"
)

type UserService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(c context.Context) ([]model.User, error) {
	users, err := s.repo.GetAllUsers(c)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}

	return users, nil
}

func (s *UserService) GetUserByID(c context.Context, id int) (*model.User, error) {
	user, err := s.repo.GetUserByID(c, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return user, nil
}
