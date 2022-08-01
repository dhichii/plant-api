package auth

import (
	"net/http"
	"plant-api/business"
	"plant-api/business/auth"
	"plant-api/utils"

	"github.com/labstack/echo/v4"
)

// Get auth API controller
type Controller struct {
	service auth.Service
}

// Construct auth API controller
func NewController(service auth.Service) *Controller {
	return &Controller{service}
}

// Controller to login
func (controller *Controller) Login(c echo.Context) error {
	loginRequest := auth.Auth{}
	c.Bind(&loginRequest)

	token, err := controller.service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		if err == business.ErrBadRequest {
			return utils.CreateWithoutDataResponse(c, http.StatusBadRequest)
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	if token == nil {
		return utils.CreateResponse(
			c,
			http.StatusUnauthorized,
			utils.Reason{Reason: "invalid email or password"},
		)
	}
	return utils.CreateResponse(c, http.StatusOK, token)
}
