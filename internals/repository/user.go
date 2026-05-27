package repository

import (
	"context"
	"microservice/internals/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
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

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	// Mock implementation, replace with actual database logic
	return &model.User{
		Name: "John Doe",
	}, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	// Mock implementation, replace with actual database logic
	return nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	// Mock implementation, replace with actual database logic
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	// Mock implementation, replace with actual database logic
	return nil
}
