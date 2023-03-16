package data

import "gorm.io/gorm"

type Reservasi struct {
	gorm.Model
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
