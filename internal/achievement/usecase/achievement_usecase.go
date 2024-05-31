package usecase

import (
	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/entity"
)

type AchievementUseCase interface {
	AddAchievement(request dto.AchievementRequestCreate) (*entity.Achievement, error)
	GetAchievements(limit, offset int) ([]entity.Achievement, int, error)
	GetAchievementById(id string) (*entity.Achievement, error)
	UpdateAchievement(request dto.AchievementRequestUpdate, id string) (*entity.Achievement, error)
	DeleteAchievement(id string) error
}
