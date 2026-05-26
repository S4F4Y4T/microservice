package repository

import (
	"context"
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

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID() (*model.User, error) {
	// Mock implementation, replace with actual database logic
	return &model.User{
		Name: "John Doe",
	}, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	// Mock implementation, replace with actual database logic
	return nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	// Mock implementation, replace with actual database logic
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	// Mock implementation, replace with actual database logic
	return nil
}
