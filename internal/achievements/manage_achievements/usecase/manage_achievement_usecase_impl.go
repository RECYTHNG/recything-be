package usecase

import (
	"mime/multipart"
	"strings"

	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/dto"
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/achievements/manage_achievements/repository"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type ManageAchievementUsecaseImpl struct {
	repository repository.ManageAchievementRepository
}

func NewManageAchievementUsecase(repository repository.ManageAchievementRepository) *ManageAchievementUsecaseImpl {
	return &ManageAchievementUsecaseImpl{repository: repository}
}

func (repository ManageAchievementUsecaseImpl) GetAllArchievementUsecase() ([]*archievement.Achievement, error) {
	achievements, err := repository.repository.GetAllArchievement()
	if err != nil {
		return nil, err
	}
	return achievements, nil
}

func (repository ManageAchievementUsecaseImpl) GetAchievementByIdUsecase(id int) (*archievement.Achievement, error) {
	achievement, err := repository.repository.GetAchievementById(id)
	if err != nil {
		return nil, pkg.ErrAchievementNotFound
	}

	return achievement, nil
}

func (repository ManageAchievementUsecaseImpl) UpdateAchievementUsecase(request *dto.UpdateAchievementRequest, badge []*multipart.FileHeader, id int) error {
	if len(badge) == 0 {
		return pkg.ErrBadge
	}
	if len(badge) > 1 {
		return pkg.ErrBadgeMaximum
	}
	validImages, errImages := helper.ImagesValidation(badge)
	if errImages != nil {
		return errImages
	}
	urlBadge, errUpload := helper.UploadToCloudinary(validImages[0], "achievements_badge")
	if errUpload != nil {
		return pkg.ErrUploadCloudinary
	}

	achievement, err := repository.repository.GetAchievementById(id)
	if err != nil {
		return pkg.ErrAchievementNotFound
	}
	achievement.Level = strings.ToLower(request.Level)
	achievement.TargetPoint = request.TargetPoint
	achievement.BadgeUrl = urlBadge
	if err := repository.repository.UpdateAchievement(achievement, id); err != nil {
		return err
	}
	return nil
}

func (repository ManageAchievementUsecaseImpl) DeleteAchievementUsecase(id int) error {
	if _, err := repository.repository.GetAchievementById(id); err != nil {
		return pkg.ErrAchievementNotFound
	}
	if err := repository.repository.DeleteAchievement(id); err != nil {
		return err
	}
	return nil
}
