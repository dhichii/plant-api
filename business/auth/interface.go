package auth

import "plant-api/api/v1/auth/response"

// Outgoing port for auth
type Service interface {
	Login(email, password string) (*response.Token, error)
}
