package main

import (
	"airbnb/app/config"
	"airbnb/app/database"
	router "airbnb/app/route"

	"github.com/labstack/echo/v4"
)

func main() {
	// docs.InitSwagger()
	cfg := config.InitConfig()
	db := database.InitDBMysql(*cfg)
	database.InitMigration(db)

	e := echo.New()
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
