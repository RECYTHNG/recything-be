package usecase

import (
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/repository"
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
		return nil, pkg.ErrAchievementLevelAlreadyExist
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

func (repository ManageAchievementUsecaseImpl) GetAchievementByIdUsecase(id int) (*archievement.Archievement, error) {
	achievement, err := repository.repository.GetAchievementById(id)
	if err != nil {
		return nil, pkg.ErrAchievementNotFound
	}

	return achievement, nil
}

func (repository ManageAchievementUsecaseImpl) UpdateAchievementUsecase(request *dto.UpdateAchievementRequest, id int) error {
	achievement, err := repository.repository.GetAchievementById(id)
	if err != nil {
		return pkg.ErrAchievementNotFound
	}
	achievement.Level = request.Level
	achievement.TargetPoint = request.TargetPoint
	achievement.BadgeUrl = request.BadgeUrl
	if err := repository.repository.UpdateAchievement(achievement, id); err != nil {
		return err
	}
	return nil
}
