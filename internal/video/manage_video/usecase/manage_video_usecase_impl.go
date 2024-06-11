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
	if len(request.VideoCategories) == 0 {
		return pkg.ErrVideoCategory
	}
	if len(request.TrashCategories) == 0 {
		return pkg.ErrVideoTrashCategory
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

	var videoCategories []video.VideoCategory
	var trashCategories []video.TrashCategory

	for _, category := range request.VideoCategories {
		name := strings.ToLower(category.Name)
		if err := usecase.manageVideoRepository.FindNameCategoryVideo(name); err == gorm.ErrRecordNotFound {
			return pkg.ErrNameCategoryVideoNotFound
		}
		videoCategory := video.VideoCategory{
			Name:      name,
			DeletedAt: gorm.DeletedAt{},
		}
		videoCategories = append(videoCategories, videoCategory)
	}
	for _, category := range request.TrashCategories {
		name := strings.ToLower(category.Name)
		if err := usecase.manageVideoRepository.FindNamaTrashCategory(name); err == gorm.ErrRecordNotFound {
			return pkg.ErrNameTrashCategoryNotFound
		}
		trashCategory := video.TrashCategory{
			Name:      name,
			DeletedAt: gorm.DeletedAt{},
		}
		trashCategories = append(trashCategories, trashCategory)
	}
	view, errGetView := helper.GetVideoViewCount(request.LinkVideo)
	if errGetView != nil {
		return errGetView
	}
	urlThumbnail, errUpload := helper.UploadToCloudinary(validImages[0], "video_thumbnail")
	if errUpload != nil {
		return pkg.ErrUploadCloudinary
	}
	intView := int(view)
	video := video.Video{
		Title:           request.Title,
		Description:     request.Description,
		Thumbnail:       urlThumbnail,
		Link:            request.LinkVideo,
		Viewer:          intView,
		VideoCategories: videoCategories,
		TrashCategories: trashCategories,
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

// func (usecase *ManageVideoUsecaseImpl) UpdateDataVideoUseCase(request *dto.UpdateDataVideoRequest, thumbnail []*multipart.FileHeader, id int) error {
// 	if len(thumbnail) > 1 {
// 		return pkg.ErrThumbnailMaximum
// 	}
// 	var urlThumbnail string
// 	if len(thumbnail) == 1 {
// 		validImages, errImages := helper.ImagesValidation(thumbnail)
// 		if errImages != nil {
// 			return errImages
// 		}
// 		urlThumbnailUpload, errUpload := helper.UploadToCloudinary(validImages[0], "video_thumbnail_update")
// 		if errUpload != nil {
// 			return pkg.ErrUploadCloudinary
// 		}
// 		urlThumbnail = urlThumbnailUpload
// 	}

// 	video, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id)
// 	if err != nil {
// 		return pkg.ErrVideoNotFound
// 	}

// 	if request.Title != "" {
// 		video.Title = request.Title
// 	}
// 	if request.Description != "" {
// 		video.Description = request.Description
// 	}
// 	if urlThumbnail != "" {
// 		video.Thumbnail = urlThumbnail
// 	}
// 	if request.LinkVideo != "" {
// 		view, errGetView := helper.GetVideoViewCount(request.LinkVideo)
// 		if errGetView != nil {
// 			return errGetView
// 		}
// 		if view != 0 {
// 			intView := int(view)
// 			video.Viewer = intView
// 		}
// 		video.Link = request.LinkVideo
// 	}
// 	if request.CategoryId != 0 {
// 		if _, err := usecase.manageVideoRepository.GetCategoryVideoById(request.CategoryId); err != nil {
// 			return pkg.ErrVideoCategoryNotFound
// 		}
// 		video.VideoCategoryID = request.CategoryId
// 	}

// 	if err := usecase.manageVideoRepository.UpdateDataVideo(video, id); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (usecase *ManageVideoUsecaseImpl) DeleteDataVideoUseCase(id int) error {
	if _, err := usecase.manageVideoRepository.GetDetailsDataVideoById(id); err != nil {
		return pkg.ErrVideoNotFound
	}
	if err := usecase.manageVideoRepository.DeleteDataVideo(id); err != nil {
		return err
	}
	return nil
}
