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

	app := server.NewEchoServer(conf, db)
	app.Start()
}
