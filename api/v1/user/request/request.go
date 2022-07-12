package request

import (
	"plant-api/business/user"
	"plant-api/utils"
)

type Request struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRequest struct {
	Name string `json:"name"`
}

func (req *Request) MapToModel() user.User {
	return user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	}
}

func (req *UpdateRequest) MapToModel() user.User {
	return user.User{
		Name: req.Name,
	}
}
