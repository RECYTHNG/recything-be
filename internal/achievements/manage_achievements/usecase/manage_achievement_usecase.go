package usecase

import (
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
)

type ManageAchievementUsecase interface {
	CreateArchievementUsecase(request *dto.CreateArchievementRequest) (*archievement.Archievement, error)
	GetAllArchievementUsecase() ([]*archievement.Archievement, error)
	GetAchievementByIdUsecase(id int) (*archievement.Archievement, error)
}
