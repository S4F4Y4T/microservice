package bootstrap

import (
	"microservice/internals/handler"
	"microservice/internals/repository"
	"microservice/internals/service"
)

type App struct {
	UserHandler *handler.UserHandler
}

func Register() *App {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	return &App{
		UserHandler: handler,
	}
}
