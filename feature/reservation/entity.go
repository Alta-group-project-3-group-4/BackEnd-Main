package reservation

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	Id         uint
	UserId     uint
	HomestayId uint `validate:"required"`
	Ticket     string
	BookDate   string
	CheckIn    string `validate:"required"`
	CheckOut   string `validate:"required"`
	Price      int    `validate:"required"`
	Status     string
	Bank       string `validate:"required"`
	VANumber   string
}

type ReservationHandler interface {
	Create() echo.HandlerFunc
}

type ReservationService interface {
	Create(newReservation Core) (Core, error)
}

type ReservationData interface {
	Create(newReservation Core) (Core, error)
}
