package usecase

import (
	"github.com/sawalreverr/recything/internal/archievements/manage_archievements/dto"
	archievement "github.com/sawalreverr/recything/internal/archievements/manage_archievements/entity"
)

type ManageAchievementUsecase interface {
	Create(request *dto.CreateArchievementRequest) (*archievement.Archievement, error)
}
