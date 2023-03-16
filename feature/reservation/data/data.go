package data

import (
	"airbnb/feature/reservation"

	"gorm.io/gorm"
)

type reservationData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.ReservationData {
	return &reservationData{
		db: db,
	}
}

func (rd *reservationData) Create(newReservation reservation.Core) (reservation.Core, error) {
	model := ToData(newReservation)
	tx := rd.db.Create(&model)
	if tx.Error != nil {
		return reservation.Core{}, tx.Error
	}

	return ToCore(model), nil
}
