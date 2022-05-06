package modules

import (
	"plant-api/api"
	uc "plant-api/api/v1/user"
	us "plant-api/business/user"
	ur "plant-api/repositories/user"

	"gorm.io/gorm"
)

// Register the modules
func RegisterModules(db *gorm.DB) api.Controller {
	// Initiate user
	userRepository := ur.NewMysqlRepository(db)
	userService := us.NewService(userRepository)
	userV1Controller := uc.NewController(userService)

	// Put all controllers together
	controllers := api.Controller{
		UserV1Controller: userV1Controller,
	}
	return controllers
}
