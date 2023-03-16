package service

import (
	"airbnb/feature/reservation"
	helper "airbnb/utils/midtrans"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type reservationSrv struct {
	qry     reservation.ReservationData
	payment helper.PaymentGateway
	vld     *validator.Validate
}

func New(rd reservation.ReservationData, p helper.PaymentGateway, vld *validator.Validate) reservation.ReservationService {
	return &reservationSrv{
		qry:     rd,
		payment: p,
		vld:     vld,
	}
}

func (bs *reservationSrv) Create(newReservation reservation.Core) (reservation.Core, error) {

	if err := bs.vld.Struct(newReservation); err != nil {
		log.Println("validation err", err)
		return reservation.Core{}, err
	}

	// Assign some default transactions
	id := newReservation.UserId
	newReservation.Ticket = fmt.Sprintf("INV-%d-%s", id, time.Now().Format("20060102-150405"))
	newReservation.BookDate = time.Now().Format("2006-01-02")
	newReservation.Status = "PENDING"

	// Charge transaction to midtrans and get the response
	vaNumber, errMidtrans := bs.payment.ChargeTransaction(newReservation.Ticket, newReservation.Price, newReservation.Bank)
	if errMidtrans != nil {
		log.Println(errMidtrans)
		return reservation.Core{}, errors.New("charge transaction failed due to internal server error, please try again")
	}
	newReservation.VANumber = vaNumber

	// Create booking
	res, err := bs.qry.Create(newReservation)
	if err != nil {
		log.Println(err.Error())
		return reservation.Core{}, errors.New("internal server error")
	}

	return res, nil
}
