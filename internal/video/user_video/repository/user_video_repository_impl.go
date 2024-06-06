package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoRepositoryImpl struct {
	DB database.Database
}

func NewUserVideoRepository(db database.Database) *UserVideoRepositoryImpl {
	return &UserVideoRepositoryImpl{DB: db}
}

func (repository *UserVideoRepositoryImpl) GetAllVideo() (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.DB.GetDB().Find(&videos).Error; err != nil {
		return nil, err
	}
	return &videos, nil
}

func (repository *UserVideoRepositoryImpl) SearchVideoByTitle(title string) (*[]video.Video, error) {
	var video []video.Video
	if err := repository.DB.GetDB().Where("title LIKE ?", "%"+title+"%").Find(&video).Error; err != nil {
		return nil, err
	}
	return &video, nil
}
