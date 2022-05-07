package modules

import (
	"plant-api/api"
	ac "plant-api/api/v1/auth"
	uc "plant-api/api/v1/user"
	as "plant-api/business/auth"
	us "plant-api/business/user"
	"plant-api/config"
	ur "plant-api/repositories/user"

	"gorm.io/gorm"
)

// Register the modules
func RegisterModules(db *gorm.DB, cfg config.Config) api.Controller {
	// Initiate user
	userRepository := ur.NewMysqlRepository(db)
	userService := us.NewService(userRepository)
	userV1Controller := uc.NewController(userService)

	// Initiate auth
	authService := as.NewService(userRepository, cfg)
	authController := ac.NewController(authService)

	// Put all controllers together
	controllers := api.Controller{
		UserV1Controller: userV1Controller,
		AuthController: authController,
	}
	return controllers
}
