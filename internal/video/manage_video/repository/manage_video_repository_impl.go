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

func (repository *ManageVideoRepositoryImpl) FindNameCategoryVideo(name string) error {
	var category video.VideoCategory
	if err := repository.DB.GetDB().Where("name = ?", name).First(&category).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) FindNamaTrashCategory(name string) error {
	var category video.TrashCategory
	if err := repository.DB.GetDB().Where("name = ?", name).First(&category).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) GetAllCategoryVideo() ([]string, error) {
	var categories []string
	if err := repository.DB.GetDB().Model(&video.VideoCategory{}).Distinct("name").Pluck("name", &categories).
		Error; err != nil {

	}
	return categories, nil
}

func (repository *ManageVideoRepositoryImpl) GetAllTrashCategoryVideo() ([]string, error) {
	var categories []string
	if err := repository.DB.GetDB().Model(&video.TrashCategory{}).Distinct("name").Pluck("name", &categories).
		Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repository *ManageVideoRepositoryImpl) GetCategoryVideoById(id int) (*video.VideoCategory, error) {
	var category video.VideoCategory
	if err := repository.DB.GetDB().Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *ManageVideoRepositoryImpl) GetAllDataVideoPagination(limit int, page int) ([]video.Video, int, error) {
	var videos []video.Video
	var total int64
	offset := (page - 1) * limit
	if err := repository.DB.GetDB().Model(&video.Video{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := repository.DB.GetDB().Limit(limit).Offset(offset).Order("id desc").Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	return videos, int(total), nil

}

func (repository *ManageVideoRepositoryImpl) GetDetailsDataVideoById(id int) (*video.Video, error) {
	var video video.Video
	if err := repository.DB.GetDB().
		Preload("VideoCategories").
		Preload("TrashCategories").
		Where("id = ?", id).
		First(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (repository *ManageVideoRepositoryImpl) UpdateDataVideo(video *video.Video, id int) error {
	if err := repository.DB.GetDB().Model(&video).Where("id = ?", id).Updates(&video).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ManageVideoRepositoryImpl) DeleteDataVideo(id int) error {
	if err := repository.DB.GetDB().Delete(&video.Video{}, id).Error; err != nil {
		return err
	}
	return nil
}
