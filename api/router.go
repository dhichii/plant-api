package api

import (
	authController "plant-api/api/v1/auth"
	userV1Controller "plant-api/api/v1/user"
	nativeV1Controller "plant-api/api/v1/native"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	UserV1Controller *userV1Controller.Controller
	AuthController   *authController.Controller
	NativeController *nativeV1Controller.Controller
}

func InitRouter(e *echo.Echo, controller Controller, jwtSecret string) {
	v1 := e.Group("/v1")

	jwtMiddleware := middleware.JWT([]byte(jwtSecret))
	// User Admin route
	v1.POST("/admin", controller.UserV1Controller.Create, jwtMiddleware)
	v1.GET("/admin", controller.UserV1Controller.GetAll, jwtMiddleware)
	v1.GET("/admin/:id", controller.UserV1Controller.Get, jwtMiddleware)
	v1.PUT("/admin/:id", controller.UserV1Controller.Update, jwtMiddleware)

	// Auth route
	v1.POST("/login", controller.AuthController.Login)

	// Native route
	v1.POST("/natives", controller.NativeController.Create)
	v1.GET("/natives", controller.NativeController.GetAll)
}
