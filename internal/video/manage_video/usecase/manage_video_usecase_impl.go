package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	repository "github.com/sawalreverr/recything/internal/video/manage_video/repository"
	"github.com/sawalreverr/recything/pkg"
)

type ManageVideoUsecaseImpl struct {
	manageVideoRepository repository.ManageVideoRepository
}

func NewManageVideoUsecaseImpl(manageVideoRepository repository.ManageVideoRepository) *ManageVideoUsecaseImpl {
	return &ManageVideoUsecaseImpl{
		manageVideoRepository: manageVideoRepository,
	}
}

func (usecase *ManageVideoUsecaseImpl) CreateDataVideoUseCase(video *video.Video) error {
	if err := usecase.manageVideoRepository.FindTitleVideo(video.Title); err == nil {
		return pkg.ErrVideoTitleAlreadyExist
	}
	if err := usecase.manageVideoRepository.CreateDataVideo(video); err != nil {
		return err
	}

	return nil
}
