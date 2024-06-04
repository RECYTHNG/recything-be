package repository

import (
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
)

type ManageAchievementRepository interface {
	CreateAchievement(achievement *archievement.Archievement) (*archievement.Archievement, error)
	FindArchievementByLevel(level string) (*archievement.Archievement, error)
	GetAllArchievement() ([]*archievement.Archievement, error)
	GetAchievementById(id int) (*archievement.Archievement, error)
	UpdateAchievement(achievement *archievement.Archievement, id int) error
	DeleteAchievement(id int) error
}
