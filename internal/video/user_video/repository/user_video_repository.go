package repository

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoRepository interface {
	GetAllVideo() (*[]video.Video, error)
	SearchVideoByKeyword(keyword string) (*[]video.Video, error)
	SearchVideoByCategoryVideo(categoryVideo string) (*[]video.Video, error)
	SearchVideoByTrashCategoryVideo(trashCategory string) (*[]video.Video, error)
	GetVideoDetail(id int) (*video.Video, *[]video.Comment, error)
	AddComment(comment *video.Comment) error
	UpdateViewer(view int, id int) error
}
