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

	// Init Custom Data
	db.InitCustomDatas()

	// Init Tasks
	db.InitTasks()

	// Init Task Steps
	db.InitTaskSteps()

	// Init Achievements
	db.InitAchievements()

	// Init About us
	db.InitAboutUs()

	// Init Waste Categories
	db.InitWasteCategories()

	// Init Content Categories
	db.InitContentCategories()

	// Init Article
	db.InitArticle()

	// Init Videos
	db.InitDataVideos()

	// Init Article
	db.InitArticle()

	app := server.NewEchoServer(conf, db)

	// cronjob for update status task
	c := cron.New()

	taskRepo := repository.NewManageTaskRepository(db)
	c.AddFunc("@daily", func() {
		log.Println("Updating task challenge status...")
		taskRepo.UpdateTaskChallengeStatus()
	})

	c.Start()
	defer c.Stop()

	app.Start()
}
