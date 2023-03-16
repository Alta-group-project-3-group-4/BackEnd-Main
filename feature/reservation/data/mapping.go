package data

import (
	"airbnb/feature/reservation"

	"gorm.io/gorm"
)

func ToData(core reservation.Core) Reservasi {
	return Reservasi{
		Model:      gorm.Model{ID: core.Id},
		UserId:     core.UserId,
		HomestayId: core.HomestayId,
		Ticket:     core.Ticket,
		BookDate:   core.BookDate,
		CheckIn:    core.CheckIn,
		CheckOut:   core.CheckOut,
		Price:      core.Price,
		Status:     core.Status,
		Bank:       core.Bank,
		VANumber:   core.VANumber,
	}
}

func ToCore(data Reservasi) reservation.Core {
	return reservation.Core{
		Id:         data.ID,
		UserId:     data.UserId,
		HomestayId: data.HomestayId,
		Ticket:     data.Ticket,
		BookDate:   data.BookDate,
		CheckIn:    data.CheckIn,
		CheckOut:   data.CheckOut,
		Price:      data.Price,
		Status:     data.Status,
		Bank:       data.Bank,
		VANumber:   data.VANumber,
	}
}
