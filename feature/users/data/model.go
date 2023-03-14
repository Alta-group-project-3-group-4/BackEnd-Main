package data

import (
	"airbnb/utils/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	HomestayId uint
	// Homestay   *homestay.Homestay `gorm:"foreignKey:HomestayId"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	About    string
	Address  string
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password, err = helpers.HashPassword(user.Password)
	return
}
