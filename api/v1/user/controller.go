package user

import (
	"net/http"
	"plant-api/api/common"
	"plant-api/api/v1/user/request"
	"plant-api/business"
	"plant-api/business/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get user API controller
type Controller struct {
	service user.Service
}

// Construct user API controller
func NewController(service user.Service) *Controller {
	return &Controller{service}
}

// Controller to create user
func (controller *Controller) Create(c echo.Context) error {
	request := request.Request{}
	c.Bind(&request)
	newUser := request.MapToModel()
	newUser.Role = "admin"
	if err := controller.service.Create(newUser); err != nil {
		if err == business.ErrConflict {
			return c.JSON(
				http.StatusConflict,
				common.ConflictResponse("email is already used"),
			)
		}
		return c.JSON(
			http.StatusInternalServerError,
			common.InternalServerErrorResponse(),
		)
	}
	return c.JSON(http.StatusCreated, common.SuccessResponseWithoutData())
}

// Controller to get all users
func (controller *Controller) GetAll(c echo.Context) error {
	users, err := controller.service.GetAll()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			common.InternalServerErrorResponse(),
		)
	}
	return c.JSON(http.StatusOK, users)
}

// Controller to get user by given id
func (controller *Controller) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.service.Get(id)
	if err != nil {
		if err == business.ErrNotFound {
			return c.JSON(http.StatusNotFound, common.NotFoundResponse())
		}
		return c.JSON(
			http.StatusInternalServerError,
			common.InternalServerErrorResponse(),
		)
	}
	return c.JSON(http.StatusOK, user)
}

// Controller to update user
func (controller *Controller) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	request := request.UpdateRequest{}
	c.Bind(&request)

	if err := controller.service.Update(id, request.MapToModel()); err != nil {
		if err == business.ErrNotFound {
			return c.JSON(http.StatusNotFound, common.NotFoundResponse())
		}
		return c.JSON(
			http.StatusInternalServerError,
			common.InternalServerErrorResponse(),
		)
	}
	return c.NoContent(http.StatusNoContent)
}
