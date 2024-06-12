package repository

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoRepository interface {
	CreateDataVideo(video *video.Video) error
	FindTitleVideo(title string) error
	FindNameCategoryVideo(name string) error
	FindNamaTrashCategory(name string) error
	GetAllCategoryVideo() ([]string, error)
	GetAllTrashCategoryVideo() ([]string, error)
	GetCategoryVideoById(id int) (*video.VideoCategory, error)
	GetAllDataVideoPagination(limit int, page int) ([]video.Video, int, error)
	GetDetailsDataVideoById(id int) (*video.Video, error)
	UpdateDataVideo(video *video.Video, id int) error
	DeleteDataVideo(id int) error
}
