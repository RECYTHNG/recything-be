package main

import (
	"github.com/sawalreverr/recything/config"
	"github.com/sawalreverr/recything/internal/database"
	"github.com/sawalreverr/recything/internal/server"
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

	app.Start()
}
