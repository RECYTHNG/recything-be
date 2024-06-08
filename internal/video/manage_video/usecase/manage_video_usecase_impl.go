package usecase

import (
	"mime/multipart"
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

func (usecase *ManageVideoUsecaseImpl) CreateDataVideoUseCase(request *dto.CreateDataVideoRequest, thumbnail []*multipart.FileHeader) error {
	if len(thumbnail) == 0 {
		return pkg.ErrThumbnail
	}
	if len(thumbnail) > 1 {
		return pkg.ErrThumbnailMaximum
	}
	validImages, errImages := helper.ImagesValidation(thumbnail)
	if errImages != nil {
		return errImages
	}

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
	urlThumbnail, errUpload := helper.UploadToCloudinary(validImages[0], "video_thumbnail_update")
	if errUpload != nil {
		return pkg.ErrUploadCloudinary
	}
	intView := int(view)
	video := video.Video{
		Title:           request.Title,
		Description:     request.Description,
		Thumbnail:       urlThumbnail,
		Link:            request.LinkVideo,
		VideoCategoryID: request.CategoryId,
		Viewer:          intView,
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

func (usecase *ManageVideoUsecaseImpl) GetAllDataVideoPaginationUseCase(limit int, page int) ([]video.Video, int, error) {
	videos, count, err := usecase.manageVideoRepository.GetAllDataVideoPagination(limit, page)
	if err != nil {
		return nil, 0, err
	}
	return videos, count, nil
}

func (usecase *ManageVideoUsecaseImpl) GetDetailsDataVideoByIdUseCase(id int) (*video.Video, error) {
	video, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrVideoNotFound
		}
		return nil, err
	}
	return video, nil
}

func (usecase *ManageVideoUsecaseImpl) UpdateDataVideoUseCase(request *dto.UpdateDataVideoRequest, thumbnail []*multipart.FileHeader, id int) error {
	if len(thumbnail) == 0 {
		return pkg.ErrThumbnail
	}
	if len(thumbnail) > 1 {
		return pkg.ErrThumbnailMaximum
	}
	validImages, errImages := helper.ImagesValidation(thumbnail)
	if errImages != nil {
		return errImages
	}

	if _, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id); err != nil {
		return pkg.ErrVideoNotFound
	}
	if _, err := usecase.manageVideoRepository.GetCategoryVideoById(request.CategoryId); err != nil {
		return pkg.ErrVideoCategoryNotFound
	}
	view, errGetView := helper.GetVideoViewCount(request.LinkVideo)
	if errGetView != nil {
		return errGetView
	}
	urlThumbnail, errUpload := helper.UploadToCloudinary(validImages[0], "video_thumbnail_update")
	if errUpload != nil {
		return pkg.ErrUploadCloudinary
	}

	intView := int(view)
	video := video.Video{
		Title:           request.Title,
		Description:     request.Description,
		Thumbnail:       urlThumbnail,
		Link:            request.LinkVideo,
		VideoCategoryID: request.CategoryId,
		Viewer:          intView,
	}
	if err := usecase.manageVideoRepository.UpdateDataVideo(&video, id); err != nil {
		return err
	}
	return nil
}

func (usecase *ManageVideoUsecaseImpl) DeleteDataVideoUseCase(id int) error {
	if _, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id); err != nil {
		return pkg.ErrVideoNotFound
	}
	if err := usecase.manageVideoRepository.DeleteDataVideo(id); err != nil {
		return err
	}
	return nil
}
