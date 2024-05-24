package database

import (
	"log"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
