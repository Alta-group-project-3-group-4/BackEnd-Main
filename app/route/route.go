package router

import (
	middlewares "airbnb/app/middleware"
	"airbnb/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	middlewares.Cors(e)
	middlewares.BasicLogger(e)
	e.GET("/", index)
	SwaggerRouter(e)
	AuthRouter(db, e)
	UserRouter(db, e)
	HomeRouter(db, e)
	CommentRouter(db, e)
	ImageRouter(db, e)
}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to API airbnb Project",
		"documentation": "/swagger/index.html",
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", data))
}
