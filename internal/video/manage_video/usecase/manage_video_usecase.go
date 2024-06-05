package usecase

import (
	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoUsecase interface {
	CreateDataVideoUseCase(video *video.Video) error
	CreateCategoryVideoUseCase(request *dto.CreateCategoryVideoRequest) error
}
