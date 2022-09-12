package api

import (
	"plant-api/api/middleware"
	accountV1Controller "plant-api/api/v1/account"
	authController "plant-api/api/v1/auth"
	nativeV1Controller "plant-api/api/v1/native"
	plantV1Controller "plant-api/api/v1/plant"
	userV1Controller "plant-api/api/v1/user"

	"github.com/labstack/echo/v4"
	eMiddleware "github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	AuthController    *authController.Controller
	AccountController *accountV1Controller.Controller
	UserV1Controller  *userV1Controller.Controller
	NativeController  *nativeV1Controller.Controller
	PlantController   *plantV1Controller.Controller
}

func InitRouter(e *echo.Echo, controller Controller, jwtSecret string) {
	v1 := e.Group("/v1")

	// Create auth JWT group
	authV1 := v1.Group("")
	authV1.Use(eMiddleware.JWT([]byte(jwtSecret)))

	// Auth route
	v1.POST("/login", controller.AuthController.Login)

	// Account route
	authV1.PATCH(
		"/accounts/email/:id",
		controller.AccountController.UpdateEmail,
		middleware.GrantByIDOrSuper,
	)
	authV1.PATCH(
		"/accounts/password/:id",
		controller.AccountController.UpdatePassword,
		middleware.GrantByIDOrSuper,
	)

	// User Admin route
	authV1.POST("/admins", controller.UserV1Controller.Create, middleware.GrantSuper)
	authV1.GET("/admins", controller.UserV1Controller.GetAll, middleware.GrantSuper)
	authV1.GET("/admins/:id", controller.UserV1Controller.Get, middleware.GrantByIDOrSuper)
	authV1.PATCH("/admins/:id", controller.UserV1Controller.Update, middleware.GrantByIDOrSuper)

	// Native route
	authV1.POST("/natives", controller.NativeController.Create, middleware.GrantAdmin)
	authV1.GET("/natives", controller.NativeController.GetAll, middleware.GrantAdmin)

	// Plant route
	authV1.POST("/plants", controller.PlantController.Create, middleware.GrantAdmin)
	authV1.PUT("/plants/:id", controller.PlantController.Update, middleware.GrantAdmin)
	authV1.DELETE("/plants/:id", controller.PlantController.Delete, middleware.GrantAdmin)
	v1.GET("/plants", controller.PlantController.GetAll)
	v1.GET("/plants/:id", controller.PlantController.GetDetail)
}
