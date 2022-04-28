package main

import (
	"fmt"
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

	// connect to mysql
	db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", cfg.Port)))
}