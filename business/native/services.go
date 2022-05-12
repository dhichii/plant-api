package native

import "plant-api/business"

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
func (s *service) GetAll() ([]Native, error) {
	return s.repository.GetAll()
}

// Get native by given id
func (s *service) Get(id int) (*Native, error) {
	native, err := s.repository.Get(id)
	if err.Error() == "record not found" {
		return nil, business.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return native, nil
}

// Get native by given name
func (s *service) GetByName(name string) (*Native, error) {
	native, err := s.repository.GetByName(name)
	if err.Error() == "record not found" {
		return nil, business.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return native, nil
}