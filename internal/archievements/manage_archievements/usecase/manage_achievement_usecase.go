package usecase

import (
	"github.com/sawalreverr/recything/internal/archievements/manage_archievements/dto"
	archievement "github.com/sawalreverr/recything/internal/archievements/manage_archievements/entity"
)

type ManageAchievementUsecase interface {
	CreateArchievementUsecase(request *dto.CreateArchievementRequest) (*archievement.Archievement, error)
	GetAllArchievementUsecase() ([]*archievement.Archievement, error)
	GetAchievementByIdUsecase(id int) (*archievement.Archievement, error)
}
