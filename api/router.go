package api

import (
	authController "plant-api/api/v1/auth"
	userV1Controller "plant-api/api/v1/user"
	nativeV1Controller "plant-api/api/v1/native"
	plantV1Controller "plant-api/api/v1/plant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	UserV1Controller *userV1Controller.Controller
	AuthController   *authController.Controller
	NativeController *nativeV1Controller.Controller
	PlantController *plantV1Controller.Controller
}

func InitRouter(e *echo.Echo, controller Controller, jwtSecret string) {
	v1 := e.Group("/v1")

	// Create auth JWT group
	authV1 := v1.Group("")
	authV1.Use(middleware.JWT([]byte(jwtSecret)))

	// User Admin route
	authV1.POST("/admin", controller.UserV1Controller.Create)
	authV1.GET("/admin", controller.UserV1Controller.GetAll)
	authV1.GET("/admin/:id", controller.UserV1Controller.Get)
	authV1.PUT("/admin/:id", controller.UserV1Controller.Update)

	// Auth route
	v1.POST("/login", controller.AuthController.Login)

	// Native route
	authV1.POST("/natives", controller.NativeController.Create)
	authV1.GET("/natives", controller.NativeController.GetAll)

	// Plant route
	authV1.POST("/plants", controller.PlantController.Create)
	authV1.PUT("/plants/:id", controller.PlantController.Update)
	authV1.DELETE("/plants/:id", controller.PlantController.Delete)
	v1.GET("/plants", controller.PlantController.GetAll)
	v1.GET("/plants/:id", controller.PlantController.GetDetail)
}
