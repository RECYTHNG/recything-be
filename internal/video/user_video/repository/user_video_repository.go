package repository

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoRepository interface {
	GetAllVideo() (*[]video.Video, error)
	SearchVideoByTitle(title string) (*[]video.Video, error)
	GetVideoDetail(id int) (*video.Video, *[]video.Comment, error)
}
