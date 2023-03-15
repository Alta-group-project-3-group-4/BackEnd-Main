package data

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId      uint
	HomeId      uint
	Rate        float32
	Description string
}
