package data

import (
	homestay "airbnb/feature/homestay/data"
	user "airbnb/feature/users/data"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId      uint
	HomeId      uint
	Rate        float32
	Description string
	User        *user.User         `gorm:"foreignKey:UserId"`
	Home        *homestay.Homestay `gorm:"foreignKey:HomeId"`
}

type HomeComment struct {
	ID          uint
	UserId      uint
	Rate        float32
	Description string
}
