package database

import (
	comment "airbnb/feature/comment/data"
	home "airbnb/feature/homestay/data"
	img "airbnb/feature/images/data"
	reservasi "airbnb/feature/reservation/data"
	user "airbnb/feature/users/data"
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&home.Homestay{},
		&comment.Comment{},
		&img.Image{},
		&reservasi.Reservasi{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
