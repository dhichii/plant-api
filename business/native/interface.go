package native

import "plant-api/api/v1/native/response"

// Ingoing port for native
type Repository interface {
	Create(*Native) (uint, error)
	GetAll() ([]response.Native, error)
}

// Outgoing port for native
type Service interface {
	Create(*Native) (uint, error)
	GetAll() ([]response.Native, error)
}
