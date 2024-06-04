package repository

import (
	archievement "github.com/sawalreverr/recything/internal/achievements/manage_achievements/entity"
	"github.com/sawalreverr/recything/internal/database"
)

type ManageAchievementRepositoryImpl struct {
	DB database.Database
}

func NewManageAchievementRepository(db database.Database) *ManageAchievementRepositoryImpl {
	return &ManageAchievementRepositoryImpl{DB: db}
}

func (repository ManageAchievementRepositoryImpl) CreateAchievement(achievement *archievement.Archievement) (*archievement.Archievement, error) {
	if err := repository.DB.GetDB().Create(achievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil
}

func (repository ManageAchievementRepositoryImpl) FindArchievementByLevel(level string) (*archievement.Archievement, error) {
	var achievement archievement.Archievement
	if err := repository.DB.GetDB().Where("level = ?", level).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (repository ManageAchievementRepositoryImpl) GetAllArchievement() ([]*archievement.Archievement, error) {
	var achievements []*archievement.Archievement
	if err := repository.DB.GetDB().Find(&achievements).Order("target_point desc").Error; err != nil {
		return nil, err
	}
	return achievements, nil
}

func (repository ManageAchievementRepositoryImpl) GetAchievementById(id int) (*archievement.Archievement, error) {
	var achievement archievement.Archievement
	if err := repository.DB.GetDB().Where("id = ?", id).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (repository ManageAchievementRepositoryImpl) UpdateAchievement(achievement *archievement.Archievement, id int) error {
	if err := repository.DB.GetDB().Where("id = ?", id).Updates(achievement).Error; err != nil {
		return err
	}
	return nil
}

func (repository ManageAchievementRepositoryImpl) DeleteAchievement(id int) error {
	if err := repository.DB.GetDB().Delete(&archievement.Archievement{}, id).Error; err != nil {
		return err
	}
	return nil
}
