package main

import (
	"log"

	"github.com/robfig/cron/v3"
	"github.com/sawalreverr/recything/config"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/internal/server"
	"github.com/sawalreverr/recything/internal/task/manage_task/repository"
)

func main() {
	conf := config.GetConfig()
	db := database.NewMySQLDatabase(conf)
	database.AutoMigrate(db)

	// Init super admin
	db.InitSuperAdmin()

	// Init Waste Materials
	db.InitWasteMaterials()

	// Init Faqs
	db.InitFaqs()

	app := server.NewEchoServer(conf, db)
	c := cron.New()

	taskRepo := repository.NewManageTaskRepository(db)
	c.AddFunc("@daily", func() {
		log.Println("Updating task challenge status...")
		repository.UpdateTaskChallengeStatus(taskRepo)
	})

	c.Start()
	defer c.Stop()

	app.Start()
}
