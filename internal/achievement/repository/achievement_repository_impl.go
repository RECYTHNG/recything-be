package repository

import (
	"github.com/sawalreverr/recything/internal/achievement/entity"
	"github.com/sawalreverr/recything/internal/database"
	"gorm.io/gorm"
)

type AchievementRepositoryImpl struct {
	DB database.Database
}

func NewAchievementRepository(db database.Database) *AchievementRepositoryImpl {
	return &AchievementRepositoryImpl{DB: db}
}

func (repository *AchievementRepositoryImpl) CreateDataAchievement(achievement *entity.Achievement) (*entity.Achievement, error) {
	if err := repository.DB.GetDB().Create(achievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil
}

func (repository *AchievementRepositoryImpl) FindAll(limit int, offset int) ([]entity.Achievement, int, error) {
	var achievements []entity.Achievement
	var totalData int64

	if err := repository.DB.GetDB().Limit(limit).Offset(offset).Find(&achievements).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}
	return achievements, int(totalData), nil
}

func (repository *AchievementRepositoryImpl) FindById(id int) (*entity.Achievement, error) {
	var achievement entity.Achievement
	if err := repository.DB.GetDB().Where("id = ?", id).First(&achievement, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &achievement, nil
}

func (repository *AchievementRepositoryImpl) Update(achievement *entity.Achievement) (*entity.Achievement, error) {
	if err := repository.DB.GetDB().Updates(achievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil
}

func (repository *AchievementRepositoryImpl) Delete(id int) error {
	if err := repository.DB.GetDB().Delete(&entity.Achievement{}, id).Error; err != nil {
		return err
	}
	return nil
}
