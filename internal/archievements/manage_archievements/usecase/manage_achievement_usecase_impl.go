package usecase

import (
	"github.com/sawalreverr/recything/internal/archievements/manage_archievements/dto"
	archievement "github.com/sawalreverr/recything/internal/archievements/manage_archievements/entity"
	"github.com/sawalreverr/recything/internal/archievements/manage_archievements/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ManageAchievementUsecaseImpl struct {
	repository repository.ManageAchievementRepository
}

func NewManageAchievementUsecase(repository repository.ManageAchievementRepository) *ManageAchievementUsecaseImpl {
	return &ManageAchievementUsecaseImpl{repository: repository}
}

func (repository ManageAchievementUsecaseImpl) CreateArchievementUsecase(request *dto.CreateArchievementRequest) (*archievement.Archievement, error) {
	findLeve, _ := repository.repository.FindArchievementByLevel(request.Level)
	if findLeve != nil {
		return nil, pkg.ErrArchievementLevelAlreadyExist
	}

	dataAchievement := &archievement.Archievement{
		Level:       request.Level,
		TargetPoint: request.TargetPoint,
		BadgeUrl:    request.BadgeUrl,
		DeletedAt:   gorm.DeletedAt{},
	}

	archievement, err := repository.repository.Create(dataAchievement)
	if err != nil {
		return nil, err
	}
	return archievement, nil

}

func (repository ManageAchievementUsecaseImpl) GetAllArchievementUsecase() ([]*archievement.Archievement, error) {
	achievements, err := repository.repository.GetAllArchievement()
	if err != nil {
		return nil, err
	}
	return achievements, nil
}
