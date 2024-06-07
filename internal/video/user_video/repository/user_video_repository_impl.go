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
	if err := repository.DB.GetDB().
		Order("created_at desc").
		Find(&videos).
		Error; err != nil {
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

	if err := repository.DB.GetDB().
		Where("id = ?", id).
		Order("created_at desc").
		First(&videos).Error; err != nil {
		return nil, nil, err
	}

	if err := repository.DB.GetDB().
		Preload("User").
		Where("video_id = ?", id).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		return nil, nil, err
	}

	return &videos, &comments, nil
}

func (repository *UserVideoRepositoryImpl) AddComment(comment *video.Comment) error {
	if err := repository.DB.GetDB().Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (repository *UserVideoRepositoryImpl) UpdateViewer(view int, id int) error {
	if err := repository.DB.GetDB().Model(&video.Video{}).Where("id = ?", id).Update("viewer", view).Error; err != nil {
		return err
	}
	return nil
}