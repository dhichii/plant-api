package main

import (
	"fmt"
	"plant-api/api"
	"plant-api/app/modules"
	"plant-api/config"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Connect to mysql
	db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	controller := modules.RegisterModules(db, cfg)
	app := echo.New()
	api.InitRouter(app, controller, cfg.JWTSecret)
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", cfg.Port)))
}
