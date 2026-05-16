package service

import (
	"microservice/internals/model"
)

type UserService struct {
	repo *model.Userepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{repo: repo}
}
