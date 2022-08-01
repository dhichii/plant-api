package plant

import "plant-api/api/v1/plant/response"

// Ingoing port for plant
type Repository interface {
	Create(*Plant) (uint, error)
	GetAll(name string) ([]response.Plant, error)
	GetDetail(id int) (*response.PlantDetail, error)
	Update(id int, plant Plant) error
	Delete(id int) error
}

// Outgoing port for plant
type Service interface {
	Create(*Plant) (uint, error)
	GetAll(name string) ([]response.Plant, error)
	GetDetail(id int) (*response.PlantDetail, error)
	Update(id int, plant Plant) error
	Delete(id int) error
}
