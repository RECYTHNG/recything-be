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

func (repository *UserVideoRepositoryImpl) GetVideoDetail(id int) (*video.Video, *[]video.Comment, error) {
	var videos video.Video
	var comments []video.Comment

	// Mencari video berdasarkan ID
	if err := repository.DB.GetDB().
		Where("id = ?", id).
		Order("created_at desc").
		First(&videos).Error; err != nil {
		return nil, nil, err
	}

	// Mencari komentar yang berhubungan dengan video berdasarkan video_id
	if err := repository.DB.GetDB().
		Where("video_id = ?", id).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		return nil, nil, err
	}

	// Kembalikan pointer ke video dan slice of comments
	return &videos, &comments, nil
}
