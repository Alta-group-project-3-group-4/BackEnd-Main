package router

import (
	middleware "airbnb/app/middleware"
	_commentData "airbnb/feature/comment/data"
	_commentHandler "airbnb/feature/comment/delivery"
	_commentService "airbnb/feature/comment/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommentRouter(db *gorm.DB, e *echo.Echo) {
	data := _commentData.New(db)
	service := _commentService.New(data)
	handler := _commentHandler.New(service)

	g := e.Group("/comment")
	g.Use(middleware.Authentication)
	g.POST("", handler.Add())
	g.DELETE("/:id", handler.Delete())
}
