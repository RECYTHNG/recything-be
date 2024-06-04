package database

import (
	"log"

	"github.com/sawalreverr/recything/internal/admin/entity"
	article "github.com/sawalreverr/recything/internal/article/entity"
	user "github.com/sawalreverr/recything/internal/user"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&user.User{},
		&entity.Admin{},
		&article.Article{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
