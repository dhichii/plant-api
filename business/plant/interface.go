package plant

import (
	"plant-api/api/v1/plant/response"
	"plant-api/business/native"
)

// Ingoing port for plant
type Repository interface {
	Create(*Plant) (uint, error)
	GetAll(name string) ([]response.Plant, error)
	GetDetail(id int) (*response.PlantDetail, error)
	GetAllNativesByPlantID(id int) ([]*response.Native, error)
	Update(id int, plant Plant) error
	UpdatePlantNatives(id int, plantNatives []*native.Native) error
	GetNativeByID(id int) error
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
