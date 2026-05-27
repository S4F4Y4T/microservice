package model

import "context"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int) error
}
