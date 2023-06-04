package main

import (
	"github.com/GroupProject2-Kelompok4/BE/app/config"
	"github.com/GroupProject2-Kelompok4/BE/app/database"
	"github.com/GroupProject2-Kelompok4/BE/app/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := database.InitDatabase(*cfg)
	database.InitMigration(db)
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
