package reservation

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	Id         uint
	UserId     uint
	HomestayId uint
	Ticket     string
	BookDate   string
	CheckIn    string
	CheckOut   string
	Price      int
	Status     string
	Bank       string
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
