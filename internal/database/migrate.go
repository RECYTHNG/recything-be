package database

import (
	"log"

	user "github.com/sawalreverr/recything/internal/user/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&user.User{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
