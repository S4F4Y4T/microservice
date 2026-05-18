package repository

import (
	"microservice/internals/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID() (*model.User, error) {
	// Mock implementation, replace with actual database logic
	return &model.User{
		Name: "John Doe",
	}, nil
}
