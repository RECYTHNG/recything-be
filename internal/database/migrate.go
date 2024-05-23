package database

import (
	"log"

	"github.com/sawalreverr/recything/internal/admin/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&entity.Admin{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
