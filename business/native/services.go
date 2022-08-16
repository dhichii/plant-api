package native

import (
	"plant-api/api/v1/native/response"
)

type service struct {
	repository Repository
}

// Construct native service object
func NewService(repo Repository) Service {
	return &service{repo}
}

// Create new native and store into database
func (s *service) Create(native *Native) (uint, error) {
	id, err := s.repository.Create(native)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get all natives
func (s *service) GetAll() ([]response.Native, error) {
	return s.repository.GetAll()
}
