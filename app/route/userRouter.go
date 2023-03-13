package router

import (
	middleware "airbnb/app/middleware"
	_userData "airbnb/feature/users/data"
	_userHandler "airbnb/feature/users/delivery"
	_userService "airbnb/feature/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB, e *echo.Echo) {
	data := _userData.New(db)
	service := _userService.New(data)
	handler := _userHandler.New(service)

	g := e.Group("/users")
	g.Use(middleware.Authentication)
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
