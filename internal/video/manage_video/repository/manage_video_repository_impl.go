package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoRepositoryImpl struct {
	DB database.Database
}

func NewManageVideoRepository(db database.Database) *ManageVideoRepositoryImpl {
	return &ManageVideoRepositoryImpl{DB: db}
}

func (repository *ManageVideoRepositoryImpl) CreateDataVideo(video *video.Video) error {
	if err := repository.DB.GetDB().Create(&video).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) FindTitleVideo(title string) error {
	var video video.Video
	if err := repository.DB.GetDB().Where("title = ?", title).First(&video).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) CreateCategoryVideo(category *video.VideoCategory) error {
	if err := repository.DB.GetDB().Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) FindNameCategoryVideo(name string) error {
	var category video.VideoCategory
	if err := repository.DB.GetDB().Where("name = ?", name).First(&category).Error; err != nil {
		return err
	}
	return nil
}
