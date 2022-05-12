package native

import "plant-api/api/v1/native/response"

// Ingoing port for native
type Repository interface {
	Create(*Native) error
	GetAll() ([]response.Native, error)
	Get(id int) (*response.Native, error)
	GetByName(string) (*Native, error)
}

// Outgoing port for native
type Service interface {
	Create(*Native) error
	GetAll() ([]response.Native, error)
	Get(id int) (*response.Native, error)
	GetByName(name string) (*Native, error)
}