package repository

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoRepository interface {
	CreateDataVideo(video *video.Video) error
	FindTitleVideo(title string) error
	CreateCategoryVideo(category *video.VideoCategory) error
	FindNameCategoryVideo(name string) error
	GetAllCategoryVideo() ([]video.VideoCategory, error)
	GetCategoryVideoById(id int) (*video.VideoCategory, error)
	GetAllDataVideoPagination(limit int, page int) ([]video.Video, int, error)
	GetDetailsDataVideoById(id int) (*video.Video, error)
}
