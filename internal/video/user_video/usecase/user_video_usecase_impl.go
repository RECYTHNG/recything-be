package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/internal/video/user_video/repository"
)

type UserVideoUsecaseImpl struct {
	Repository repository.UserVideoRepository
}

func NewUserVideoUsecase(repository repository.UserVideoRepository) *UserVideoUsecaseImpl {
	return &UserVideoUsecaseImpl{Repository: repository}
}

func (usecase *UserVideoUsecaseImpl) GetAllVideo() (*[]video.Video, error) {
	videos, err := usecase.Repository.GetAllVideo()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
