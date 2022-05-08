package user

import (
	"plant-api/business"
	"strings"
)

type service struct {
	repository Repository
}

// Construct user service object
func NewService(repo Repository) Service {
	return &service{repo}
}

/*
Create new user and store into database
will return ErrConflict if email is already used
*/
func (s *service) Create(user User) error {
	if err := s.repository.Create(user); err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			return business.ErrConflict
		}
		return err
	}
	return nil
}

// Get all users
func (s *service) GetAll() ([]User, error) {
	return s.repository.GetAll()
}

// Get user by given id
func (s *service) Get(id int) (*User, error) {
	user, err := s.repository.Get(id)
	if err.Error() == "record not found" {
		return nil, business.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

/*
Update existing user in database
will return ErrNotFound when user is not exist
*/
func (s *service) Update(id int, user User) error {
	_, err := s.repository.Get(id)
	if err.Error() == "record not found" {
		return business.ErrNotFound
	}
	if err != nil {
		return err
	}
	updatedUser := Modify(user.Name, user.Email, user.Password)
	return s.repository.Update(id, updatedUser)
}
