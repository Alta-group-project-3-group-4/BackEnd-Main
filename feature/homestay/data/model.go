package data

import (
	user "airbnb/feature/users/data"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	UserId      uint
	User        *user.User `gorm:"foreignKey:UserId"`
	Name        string
	Price       float64
	Description string
	// Gambar      []Gambar
	// Reservation []Reservation
}
