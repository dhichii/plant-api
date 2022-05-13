package plant

import (
	"net/http"
	"plant-api/api/common"
	"plant-api/business/plant"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get plant API controller
type Controller struct {
	service plant.Service
}

// Construct plant API controller
func NewController(service plant.Service) *Controller {
	return &Controller{service}
}

// Controller to create plant
func (controller *Controller) Create(c echo.Context) error {
	newPlant := &plant.Plant{}
	c.Bind(newPlant)
	if err := controller.service.Create(newPlant); err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusCreated, common.SuccessResponseWithoutData())
}

// Controller to get all plant
func (controller *Controller) GetAll(c echo.Context) error {
	plant, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, plant)
}

// Controller to get all plant by given name
func (controller *Controller) GetByName(c echo.Context) error {
	name := c.QueryParam("name")
	plant, err := controller.service.GetByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse())
	}
	return c.JSON(http.StatusOK, plant)
}

// Controller to get plant detail
func (controller *Controller) GetDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	plant, err := controller.service.GetDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, plant)
}

// Controller to update plant
func (controller *Controller) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	plant := plant.Plant{}
	c.Bind(&plant)
	if err := controller.service.Update(id, plant); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, common.SuccessResponseWithoutData())
}

// Controller to delete plant
func (controller *Controller) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.service.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, common.SuccessResponseWithoutData())
}
