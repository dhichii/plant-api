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
func (s *service) Create(native *Native) error {
	if err := s.repository.Create(native); err != nil {
		return err
	}
	return nil
}

// Get all natives
func (s *service) GetAll() ([]response.Native, error) {
	return s.repository.GetAll()
}

// Get native by given name
func (s *service) GetByName(name string) (*Native, error) {
	return s.repository.GetByName(name)
}