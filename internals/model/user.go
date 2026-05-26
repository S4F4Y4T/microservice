package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
	GetAllUsers() ([]User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
}
