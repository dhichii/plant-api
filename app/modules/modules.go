package modules

import (
	"plant-api/api"
	ac "plant-api/api/v1/auth"
	uc "plant-api/api/v1/user"
	nc "plant-api/api/v1/native"
	pc "plant-api/api/v1/plant"
	as "plant-api/business/auth"
	us "plant-api/business/user"
	pns "plant-api/business/plant_native"
	ns "plant-api/business/native"
	ps "plant-api/business/plant"
	"plant-api/config"
	ur "plant-api/repositories/user"
	nr "plant-api/repositories/native"
	pr "plant-api/repositories/plant"
	pnr "plant-api/repositories/plant_native"

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

	// Initiate native
	nativeRepo := nr.NewMysqlRepository(db)
	nativeService := ns.NewService(nativeRepo)
	nativeController := nc.NewController(nativeService)

	// Initiate plant native
	pNativeRepo := pnr.NewMysqlRepository(db)
	pNativeService := pns.NewService(pNativeRepo)

	// Initiate plant
	plantRepo := pr.NewMysqlRepository(db)
	plantService := ps.NewService(plantRepo, nativeService, pNativeService)
	plantController := pc.NewController(plantService)

	// Put all controllers together
	controllers := api.Controller{
		UserV1Controller: userV1Controller,
		AuthController: authController,
		NativeController: nativeController,
		PlantController: plantController,
	}
	return controllers
}
