package database

import (
	"log"

	"github.com/sawalreverr/recything/internal/admin/entity"
	user "github.com/sawalreverr/recything/internal/user"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&user.User{},
		&entity.Admin{},
		&entity.Article{},
		&entity.Achievement{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
