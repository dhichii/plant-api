package native

import (
	"net/http"
	"plant-api/business/native"
	"plant-api/utils"

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
	id, err := controller.service.Create(newNative)
	if err != nil {
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return utils.CreateResponse(c, http.StatusCreated, utils.CreatedResponse{ID: id})
}

// Controller to get all native
func (controller *Controller) GetAll(c echo.Context) error {
	natives, err := controller.service.GetAll()
	if err != nil {
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return utils.CreateResponse(c, http.StatusOK, natives)
}
