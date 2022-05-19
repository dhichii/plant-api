package user

import "plant-api/api/v1/user/response"

// Ingoing port for user
type Repository interface {
	Create(user User) error
	GetAll() ([]response.User, error)
	Get(id int) (*response.User, error)
	GetByEmail(email string) (*User, error)
	Update(id int, user User) error
}

// Outgoing port for user
type Service interface {
	Create(user User) error
	GetAll() ([]response.User, error)
	Get(id int) (*response.User, error)
	Update(id int, user User) error
}