package database

import (
	home "airbnb/feature/homestay/data"
	comment "airbnb/feature/comment/data"
	user "airbnb/feature/users/data"
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&home.Homestay{},
		&comment.Comment{},

	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
