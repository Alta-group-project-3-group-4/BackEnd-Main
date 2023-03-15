package database

import (
	comment "airbnb/feature/comment/data"
	user "airbnb/feature/users/data"
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&comment.Comment{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
