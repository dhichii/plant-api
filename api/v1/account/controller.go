package account

import (
	"net/http"
	"plant-api/api/v1/account/request"
	"plant-api/business/account"
	"plant-api/utils"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Get account API controller
type Controller struct {
	service account.Service
}

// Construct account API controller
func NewController(service account.Service) *Controller {
	return &Controller{service}
}

// Controller to update user email
func (controller *Controller) UpdateEmail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	request := request.EmailRequest{}
	c.Bind(&request)
	if request.Email == "" {
		return utils.CreateWithoutDataResponse(c, http.StatusBadRequest)
	}
	if err := controller.service.UpdateEmail(id, request.Email); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateWithoutDataResponse(c, http.StatusNotFound)
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

// Controller to update user password
func (controller *Controller) UpdatePassword(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	request := request.PasswordRequest{}
	c.Bind(&request)
	if request.Password == "" {
		return utils.CreateWithoutDataResponse(c, http.StatusBadRequest)
	}
	if err := controller.service.UpdatePassword(id, request.HashPassword()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateWithoutDataResponse(c, http.StatusNotFound)
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
