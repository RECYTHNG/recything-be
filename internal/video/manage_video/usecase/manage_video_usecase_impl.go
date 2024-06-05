package usecase

import (
	"strings"

	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/video/manage_video/dto"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	repository "github.com/sawalreverr/recything/internal/video/manage_video/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ManageVideoUsecaseImpl struct {
	manageVideoRepository repository.ManageVideoRepository
}

func NewManageVideoUsecaseImpl(manageVideoRepository repository.ManageVideoRepository) *ManageVideoUsecaseImpl {
	return &ManageVideoUsecaseImpl{
		manageVideoRepository: manageVideoRepository,
	}
}

func (usecase *ManageVideoUsecaseImpl) CreateDataVideoUseCase(request *dto.CreateDataVideoRequest) error {
	if err := usecase.manageVideoRepository.FindTitleVideo(request.Title); err == nil {
		return pkg.ErrVideoTitleAlreadyExist
	}
	if _, err := usecase.manageVideoRepository.GetCategoryVideoById(request.CategoryId); err != nil {
		return pkg.ErrVideoCategoryNotFound
	}
	view, errGetView := helper.GetVideoViewCount(request.LinkVideo)
	if errGetView != nil {
		return errGetView
	}
	intView := int(view)
	video := video.Video{
		Title:           request.Title,
		Description:     request.Description,
		Thumbnail:       request.UrlThumbnail,
		Link:            request.LinkVideo,
		VideoCategoryID: request.CategoryId,
		View:            intView,
		DeletedAt:       gorm.DeletedAt{},
	}
	if err := usecase.manageVideoRepository.CreateDataVideo(&video); err != nil {
		return err
	}
	return nil
}

func (usecase *ManageVideoUsecaseImpl) CreateCategoryVideoUseCase(request *dto.CreateCategoryVideoRequest) error {
	if err := usecase.manageVideoRepository.FindNameCategoryVideo(request.Name); err == nil {
		return pkg.ErrVideoCategoryNameAlreadyExist
	}
	name := strings.ToLower(request.Name)
	category := video.VideoCategory{
		Name:      name,
		DeletedAt: gorm.DeletedAt{},
	}
	if err := usecase.manageVideoRepository.CreateCategoryVideo(&category); err != nil {
		return err
	}
	return nil
}

func (usecase *ManageVideoUsecaseImpl) GetAllCategoryVideoUseCase() ([]video.VideoCategory, error) {
	categories, err := usecase.manageVideoRepository.GetAllCategoryVideo()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
