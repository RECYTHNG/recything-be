package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoUsecase interface {
	GetAllVideo() (*[]video.Video, error)
}
