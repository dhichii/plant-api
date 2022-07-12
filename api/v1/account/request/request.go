package request

import "plant-api/utils"

type EmailRequest struct {
	Email string `json:"email"`
}

type PasswordRequest struct {
	Password string `json:"password"`
}

func (req *PasswordRequest) HashPassword() string {
	return utils.HashPassword(req.Password)
}
