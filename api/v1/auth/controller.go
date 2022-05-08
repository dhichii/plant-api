package auth

import (
	"net/http"
	"plant-api/api/common"
	"plant-api/business"
	"plant-api/business/auth"

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
			return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
		}
		return c.JSON(
			http.StatusInternalServerError,
			common.InternalServerErrorResponse(),
		)
	}
	if token == "" {
		return c.JSON(
			http.StatusUnauthorized,
			common.UnauthorizedResponse("Invalid email or password"))
	}
	return c.JSON(http.StatusOK, token)
}
