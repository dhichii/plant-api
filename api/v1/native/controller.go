package native

import (
	"net/http"
	"plant-api/api/common"
	"plant-api/business/native"

	"github.com/labstack/echo/v4"
)

// Get native API controller
type Controller struct {
	service native.Service
}

// Construct native API controller
func NewController(service native.Service) *Controller {
	return &Controller{service}
}

// Controller to create native
func (controller *Controller) Create(c echo.Context) error {
	newNative := &native.Native{}
	c.Bind(&newNative)
	if err := controller.service.Create(newNative); err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusCreated, common.SuccessResponseWithoutData())
}

// Controller to get all native
func (controller *Controller) GetAll(c echo.Context) error {
	natives, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, natives)
}
