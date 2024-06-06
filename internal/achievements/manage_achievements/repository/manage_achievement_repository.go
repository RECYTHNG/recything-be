package repository

import (
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
)

type ManageAchievementRepository interface {
	CreateAchievement(achievement *archievement.Achievement) (*archievement.Achievement, error)
	FindArchievementByLevel(level string) (*archievement.Achievement, error)
	GetAllArchievement() ([]*archievement.Achievement, error)
	GetAchievementById(id int) (*archievement.Achievement, error)
	UpdateAchievement(achievement *archievement.Achievement, id int) error
	DeleteAchievement(id int) error
}
