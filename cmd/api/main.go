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

	// Init User
	db.InitUser()

	// Init super admin
	db.InitSuperAdmin()

	// Init Waste Materials
	db.InitWasteMaterials()

	// Init Waste Categories
	db.InitWasteCategories()

	// Init Content Categories
	db.InitContentCategories()

	// Init Achievements
	db.InitAchievements()

	// Init About us
	db.InitAboutUs()

	// Init Faqs
	db.InitFaqs()

	// Init Custom Data
	db.InitCustomDatas()

	// Init Tasks
	db.InitTasks()

	// Init Videos
	db.InitVideos()

	// Init Article
	db.InitArticle()

	// Init Report
	db.InitReport()

	// Init Comment
	db.InitComment()

	app := server.NewEchoServer(conf, db)

	app.Start()
}
