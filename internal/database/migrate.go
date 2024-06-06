package database

import (
	"log"

	achievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/admin/entity"
	customdata "github.com/sawalreverr/recything/internal/custom-data"
	"github.com/sawalreverr/recything/internal/faq"
	"github.com/sawalreverr/recything/internal/report"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
	user "github.com/sawalreverr/recything/internal/user"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&user.User{},
		&entity.Admin{},
		&report.Report{},
		&report.WasteMaterial{},
		&report.ReportWasteMaterial{},
		&report.ReportImage{},
		&faq.FAQ{},
		&task.TaskChallenge{},
		&task.TaskStep{},
		&user_task.UserTaskChallenge{},
		&user_task.UserTaskImage{},
		&achievement.Achievement{},
		&customdata.CustomData{},
		&video.Video{},
		&video.VideoCategory{},
		&video.Comment{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
