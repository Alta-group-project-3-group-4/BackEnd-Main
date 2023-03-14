package router

import (
	middleware "airbnb/app/middleware"
	_authData "airbnb/feature/auth/data"
	_authHandler "airbnb/feature/auth/delivery"
	_authService "airbnb/feature/auth/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB, e *echo.Echo) {
	data := _authData.New(db)
	service := _authService.New(data)
	handler := _authHandler.New(service)

	g := e.Group("/auth")
	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
	// g.POST("/forget-password", handler.Create)

	g.Use(middleware.Authentication)
	g.GET("/users", handler.GetUserLogin)
	g.POST("/change-password", handler.ChangePassword)
}
