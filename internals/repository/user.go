package repository

import "microservice/internals/model"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	// Mock implementation, replace with actual database logic
	return &model.User{
		ID:   id,
		Name: "John Doe",
	}, nil
}
