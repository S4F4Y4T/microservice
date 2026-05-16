package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Userepository interface {
	GetUserByID(id int) (*User, error)
}
