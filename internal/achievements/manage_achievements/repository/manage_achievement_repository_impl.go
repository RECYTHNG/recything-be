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

func (m ManageAchievementRepositoryImpl) Create(achievement *archievement.Archievement) (*archievement.Archievement, error) {
	if err := m.DB.GetDB().Create(achievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil
}

func (m ManageAchievementRepositoryImpl) FindArchievementByLevel(level string) (*archievement.Archievement, error) {
	var achievement archievement.Archievement
	if err := m.DB.GetDB().Where("level = ?", level).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (m ManageAchievementRepositoryImpl) GetAllArchievement() ([]*archievement.Archievement, error) {
	var achievements []*archievement.Archievement
	if err := m.DB.GetDB().Find(&achievements).Error; err != nil {
		return nil, err
	}
	return achievements, nil
}

func (m ManageAchievementRepositoryImpl) GetAchievementById(id int) (*archievement.Archievement, error) {
	var achievement archievement.Archievement
	if err := m.DB.GetDB().Where("id = ?", id).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}
