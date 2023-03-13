package database

import (
	"fmt"
	"os/user"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
