package router

import (
	middleware "airbnb/app/middleware"
	_Data "airbnb/feature/images/data"
	_Handler "airbnb/feature/images/delivery"
	_Service "airbnb/feature/images/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ImageRouter(db *gorm.DB, e *echo.Echo) {
	data := _Data.New(db)
	service := _Service.NewGambar(data)
	handler := _Handler.New(service)

	g := e.Group("/images")
	g.Use(middleware.Authentication)
	g.GET("/gambar", handler.Getall)

	g.POST("/gambar", handler.CreateGambar)

}
