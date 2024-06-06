package usecase

import (
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
)

type ManageAchievementUsecase interface {
	CreateArchievementUsecase(request *dto.CreateArchievementRequest) (*archievement.Achievement, error)
	GetAllArchievementUsecase() ([]*archievement.Achievement, error)
	GetAchievementByIdUsecase(id int) (*archievement.Achievement, error)
	UpdateAchievementUsecase(request *dto.UpdateAchievementRequest, id int) error
	DeleteAchievementUsecase(id int) error
}
