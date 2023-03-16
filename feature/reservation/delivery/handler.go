package delivery

import (
	middlewares "airbnb/app/middleware"
	"airbnb/feature/reservation"
	"airbnb/utils/helpers"
	"log"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type reservationController struct {
	srv reservation.ReservationService
}

func New(rs reservation.ReservationService) reservation.ReservationHandler {
	return &reservationController{
		srv: rs,
	}
}

func (rc *reservationController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		rr := ReservationRequest{}
		if err := c.Bind(&rr); err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
		}

		claim := middlewares.ClaimsToken(c)
		userId := claim.Id
		rr.UserId = uint(userId)

		newReservation := reservation.Core{}
		copier.Copy(&newReservation, &rr)
		if err := copier.Copy(&newReservation, &rr); err != nil {
			log.Println("create booking", err)
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
		}

		res, err := rc.srv.Create(newReservation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
		}

		response := struct {
			BookingID uint `json:"booking_id"`
		}{
			BookingID: res.Id,
		}

		return c.JSON(http.StatusOK, helpers.ResponseSuccess("success create reservation", response))
	}
}
