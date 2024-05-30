package database

import (
	"log"

	"github.com/sawalreverr/recything/internal/admin/entity"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	user "github.com/sawalreverr/recything/internal/user"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&user.User{},
		&entity.Admin{},
		task.TaskChallenge{},
		task.TaskStep{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
