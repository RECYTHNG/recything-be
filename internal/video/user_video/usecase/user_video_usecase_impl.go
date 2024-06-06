package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
	"github.com/sawalreverr/recything/internal/video/user_video/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type UserVideoUsecaseImpl struct {
	Repository repository.UserVideoRepository
}

func NewUserVideoUsecase(repository repository.UserVideoRepository) *UserVideoUsecaseImpl {
	return &UserVideoUsecaseImpl{Repository: repository}
}

func (usecase *UserVideoUsecaseImpl) GetAllVideoUsecase() (*[]video.Video, error) {
	videos, err := usecase.Repository.GetAllVideo()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (usecase *UserVideoUsecaseImpl) SearchVideoByTitleUsecase(title string) (*[]video.Video, error) {
	videos, err := usecase.Repository.SearchVideoByTitle(title)
	if err != nil {
		return nil, err
	}
	if videos == nil {
		return nil, pkg.ErrVideoNotFound
	}
	return videos, nil
}

func (usecase *UserVideoUsecaseImpl) GetVideoDetailUsecase(id int) (*video.Video, *[]video.Comment, error) {
	video, comments, err := usecase.Repository.GetVideoDetail(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, pkg.ErrVideoNotFound
		}
		return nil, nil, err
	}
	return video, comments, nil
}

func (usecase *UserVideoUsecaseImpl) AddCommentUsecase(request *dto.AddCommentRequest, userId string) error {
	if _, _, err := usecase.Repository.GetVideoDetail(request.VideoID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return pkg.ErrVideoNotFound
		}
		return err
	}

	comment := video.Comment{
		Comment: request.Comment,
		UserID:  userId,
		VideoID: request.VideoID,
	}
	if err := usecase.Repository.AddComment(&comment); err != nil {
		return err
	}
	return nil
}
