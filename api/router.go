package api

import (
	userV1Controller "plant-api/api/v1/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserV1Controller *userV1Controller.Controller
}

func InitRouter(e *echo.Echo, controller Controller) {
	v1 := e.Group("/v1")
	// User Admin route
	v1.POST("/admin", controller.UserV1Controller.Create)
	v1.GET("/admin", controller.UserV1Controller.GetAll)
	v1.GET("/admin/:id", controller.UserV1Controller.Get)
	v1.PUT("/admin/:id", controller.UserV1Controller.Update)
}
