package plant

import (
	"fmt"
	"net/http"
	"plant-api/api/v1/plant/request"
	"plant-api/business"
	"plant-api/business/plant"
	"plant-api/utils"
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
	newPlant := request.Request{}
	if err := c.Bind(&newPlant); err != nil {
		return utils.CreateWithoutDataResponse(c, http.StatusBadRequest)
	}
	id, err := controller.service.Create(newPlant.MapToModel())
	if err != nil {
		if err == business.ErrNotFound {
			return utils.CreateResponse(c, http.StatusNotFound, fmt.Sprintf("native with id %d not found", id))
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return utils.CreateResponse(c, http.StatusCreated, utils.CreatedResponse{ID: id})
}

// Controller to get all plant by given name from query
func (controller *Controller) GetAll(c echo.Context) error {
	name := c.QueryParam("name")
	plant, err := controller.service.GetAll(name)
	if err != nil {
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return utils.CreateResponse(c, http.StatusOK, plant)
}

// Controller to get plant detail
func (controller *Controller) GetDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	plant, err := controller.service.GetDetail(id)
	if err != nil {
		if err == business.ErrNotFound {
			return utils.CreateWithoutDataResponse(c, http.StatusNotFound)
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return utils.CreateResponse(c, http.StatusOK, plant)
}

// Controller to update plant
func (controller *Controller) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	existingPlant := request.Request{}
	if err := c.Bind(&existingPlant); err != nil {
		return utils.CreateWithoutDataResponse(c, http.StatusBadRequest)
	}
	if err := controller.service.Update(id, existingPlant.MapToModel()); err != nil {
		if err == business.ErrNotFound {
			return utils.CreateWithoutDataResponse(c, http.StatusNotFound)
		}
		if err.Error() == "native not found" {
			return utils.CreateResponse(
				c,
				http.StatusNotFound,
				utils.ErrorResponse{Reason: err.Error()},
			)
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

// Controller to delete plant
func (controller *Controller) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.service.Delete(id); err != nil {
		if err == business.ErrNotFound {
			return utils.CreateWithoutDataResponse(c, http.StatusNotFound)
		}
		return utils.CreateWithoutDataResponse(c, http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
