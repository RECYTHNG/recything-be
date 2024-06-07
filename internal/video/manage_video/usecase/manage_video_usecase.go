package usecase

import (
	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type ManageVideoUsecase interface {
	CreateDataVideoUseCase(request *dto.CreateDataVideoRequest) error
	CreateCategoryVideoUseCase(request *dto.CreateCategoryVideoRequest) error
	GetAllCategoryVideoUseCase() ([]video.VideoCategory, error)
	GetAllDataVideoPaginationUseCase(limit int, page int) ([]video.Video, int, error)
	GetDetailsDataVideoByIdUseCase(id int) (*video.Video, error)
	UpdateDataVideoUseCase(request *dto.UpdateDataVideoRequest, id int) error
	DeleteDataVideoUseCase(id int) error
}