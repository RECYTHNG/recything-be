package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoUsecase interface {
	CreateDataVideoUseCase(video *video.Video) error
}
