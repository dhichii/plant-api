package request

import (
	"plant-api/business/user"
	"plant-api/utils"
)

type Request struct {
	Name     string `json:"string"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *Request) MapToModel() user.User {
	return user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	}
}
