package auth

import (
	"plant-api/api/middleware"
	"plant-api/api/v1/auth/response"
	"plant-api/business"
	"plant-api/business/user"
	"plant-api/config"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo user.Repository
	cfg  config.Config
}

// Construct auth service object
func NewService(repo user.Repository, cfg config.Config) Service {
	return &service{
		repo,
		cfg,
	}
}

// Login by given user email and password, return empty if not found
func (s *service) Login(email, password string) (*response.Token, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}
	if user != nil {
		matchPassword := matchPassword(user.Password, []byte(password))
		if user.Email == email && matchPassword {
			token, err := middleware.GenerateJWT(int(user.ID), user.Role, s.cfg.JWTSecret)
			if err != nil {
				return nil, business.ErrBadRequest
			}
			tokenResponse := response.Token{Token: token}
			return &tokenResponse, nil
		}
	}
	return nil, nil
}
