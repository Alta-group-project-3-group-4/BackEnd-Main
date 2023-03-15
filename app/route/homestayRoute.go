package router

import (
	middleware "airbnb/app/middleware"
	_homeData "airbnb/feature/homestay/data"
	_homeHandler "airbnb/feature/homestay/delivery"
	_homeService "airbnb/feature/homestay/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HomeRouter(db *gorm.DB, e *echo.Echo) {
	data := _homeData.New(db)
	service := _homeService.NewRoom(data)
	handler := _homeHandler.New(service)

	g := e.Group("/homestay")
	g.Use(middleware.Authentication)
	g.GET("", handler.GetAll())
	g.GET("/:id", handler.Profile())
	g.POST("", handler.CreateRoom)
	// g.PUT("/:id", handler.)
	g.DELETE("/:id", handler.Delete())
}
