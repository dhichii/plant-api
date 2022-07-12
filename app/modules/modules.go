package modules

import (
	"plant-api/api"
	authController "plant-api/api/v1/auth"
	nativeController "plant-api/api/v1/native"
	plantController "plant-api/api/v1/plant"
	userController "plant-api/api/v1/user"
	authService "plant-api/business/auth"
	nativeService "plant-api/business/native"
	plantService "plant-api/business/plant"
	plantNativeService "plant-api/business/plant_native"
	userService "plant-api/business/user"
	"plant-api/config"
	nativeRepository "plant-api/repositories/native"
	plantRepository "plant-api/repositories/plant"
	plantNativeRepository "plant-api/repositories/plant_native"
	userRepository "plant-api/repositories/user"

	"gorm.io/gorm"
)

// Register the modules
func RegisterModules(db *gorm.DB, cfg config.Config) api.Controller {
	// Initiate user
	userRepo := userRepository.NewMysqlRepository(db)
	userService := userService.NewService(userRepo)
	userController := userController.NewController(userService)

	// Initiate auth
	authService := authService.NewService(userRepo, cfg)
	authController := authController.NewController(authService)

	// Initiate native
	nativeRepo := nativeRepository.NewMysqlRepository(db)
	nativeService := nativeService.NewService(nativeRepo)
	nativeController := nativeController.NewController(nativeService)

	// Initiate plant native
	pNativeRepo := plantNativeRepository.NewMysqlRepository(db)
	pNativeService := plantNativeService.NewService(pNativeRepo)

	// Initiate plant
	plantRepo := plantRepository.NewMysqlRepository(db)
	plantService := plantService.NewService(plantRepo, nativeService, pNativeService)
	plantController := plantController.NewController(plantService)

	// Put all controllers together
	controllers := api.Controller{
		UserV1Controller: userController,
		AuthController:   authController,
		NativeController: nativeController,
		PlantController:  plantController,
	}
	return controllers
}
