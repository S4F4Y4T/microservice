package handler

import (
	"microservice/internals/service"
	"microservice/pkg/response"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.GetAllUsers(r.Context())

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	response.Success(w, http.StatusOK, "Users retrieved successfully", user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Implement logic to handle the request and call the service method
}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Implement logic to handle the request and call the service method
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implement logic to handle the request and call the service method
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implement logic to handle the request and call the service method
}
