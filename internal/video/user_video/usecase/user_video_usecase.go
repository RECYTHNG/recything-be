package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoUsecase interface {
	GetAllVideoUsecase() (*[]video.Video, error)
	SearchVideoByTitleUsecase(title string) (*[]video.Video, error)
}
