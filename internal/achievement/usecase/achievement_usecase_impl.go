package usecase

import (
	"strconv"

	"github.com/sawalreverr/recything/internal/article/dto"
	"github.com/sawalreverr/recything/internal/article/entity"
	"github.com/sawalreverr/recything/internal/article/repository"
	"github.com/sawalreverr/recything/pkg"
)

type AchievementUseCaseImpl struct {
	repo repository.AchievementRepository
}

func NewAchievementUseCase(repo repository.AchievementRepository) AchievementUseCase {
	return &AchievementUseCaseImpl{repo: repo}
}

func (u *AchievementUseCaseImpl) AddAchievement(request dto.AchievementRequestCreate) (*entity.Achievement, error) {
	achievement := &entity.Achievement{
		Level:       request.Level,
		Lencana:     request.Lencana,
		TargetPoint: request.TargetPoint,
	}
	return u.repo.Create(achievement)
}

func (u *AchievementUseCaseImpl) GetAchievements(limit, offset int) ([]entity.Achievement, int, error) {
	return u.repo.FindAll(limit, offset)
}

func (u *AchievementUseCaseImpl) GetAchievementById(id string) (*entity.Achievement, error) {
	achievementId, err := strconv.Atoi(id)
	if err != nil {
		return nil, pkg.ErrInvalidID
	}
	return u.repo.FindById(achievementId)
}

func (u *AchievementUseCaseImpl) UpdateAchievement(request dto.AchievementRequestUpdate, id string) (*entity.Achievement, error) {
	achievementId, err := strconv.Atoi(id)
	if err != nil {
		return nil, pkg.ErrInvalidID
	}

	achievement, err := u.repo.FindById(achievementId)
	if err != nil {
		return nil, err
	}

	achievement.Level = request.Level
	achievement.Lencana = request.Lencana
	achievement.TargetPoint = request.TargetPoint

	return u.repo.Update(achievement)
}

func (u *AchievementUseCaseImpl) DeleteAchievement(id string) error {
	achievementId, err := strconv.Atoi(id)
	if err != nil {
		return pkg.ErrInvalidID
	}

	return u.repo.Delete(achievementId)
}
