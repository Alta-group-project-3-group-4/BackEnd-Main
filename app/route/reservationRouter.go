package router

import (
	"airbnb/app/config"
	middleware "airbnb/app/middleware"
	_reservationData "airbnb/feature/reservation/data"
	_reservationHandler "airbnb/feature/reservation/delivery"
	_reservationService "airbnb/feature/reservation/service"
	helper "airbnb/utils/midtrans"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ReservationRouter(db *gorm.DB, e *echo.Echo) {
	cfg := config.InitConfig()
	v := validator.New()
	coreapiMidtrans := helper.NewCoreMidtrans(cfg)

	data := _reservationData.New(db)
	service := _reservationService.New(data, coreapiMidtrans, v)
	handler := _reservationHandler.New(service)

	g := e.Group("/reservation")
	g.Use(middleware.Authentication)
	g.POST("", handler.Create())

}
