package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/pkg"
)

type UserVideoRepositoryImpl struct {
	DB database.Database
}

func NewUserVideoRepository(db database.Database) *UserVideoRepositoryImpl {
	return &UserVideoRepositoryImpl{DB: db}
}

func (repository *UserVideoRepositoryImpl) GetAllVideo() (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.DB.GetDB().
		Order("created_at desc").
		Find(&videos).
		Error; err != nil {
		return nil, err
	}
	return &videos, nil
}

func (repository *UserVideoRepositoryImpl) SearchVideoByKeyword(keyword string) (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.DB.GetDB().
		Order("created_at desc").
		Preload("VideoCategories").
		Preload("TrashCategories").
		Joins("JOIN video_categories ON video_categories.video_id = videos.id").
		Joins("JOIN trash_categories ON trash_categories.video_id = videos.id").
		Where("videos.title LIKE ? OR videos.description LIKE ? OR trash_categories.name LIKE ? OR video_categories.name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Find(&videos).Error; err != nil {
		return nil, err
	}
	videoMap := make(map[int]video.Video)
	for _, v := range videos {
		videoMap[v.ID] = v
	}

	uniqueVideos := make([]video.Video, 0, len(videoMap))
	for _, v := range videoMap {
		uniqueVideos = append(uniqueVideos, v)
	}

	return &uniqueVideos, nil
}

func (repository *UserVideoRepositoryImpl) SearchVideoByCategory(categoryType string, name string) (*[]video.Video, error) {
	var videos []video.Video
	if categoryType == "content" {
		if err := repository.DB.GetDB().
			Order("created_at desc").
			Joins("JOIN video_categories ON video_categories.video_id = videos.id").
			Where("video_categories.name LIKE ?", "%"+name+"%").
			Preload("VideoCategories").
			Preload("TrashCategories").
			Find(&videos).Error; err != nil {
			return nil, err
		}

	} else if categoryType == "waste" {
		if err := repository.DB.GetDB().
			Order("created_at desc").
			Joins("JOIN trash_categories ON trash_categories.video_id = videos.id").
			Where("trash_categories.name LIKE ?", "%"+name+"%").
			Preload("VideoCategories").
			Preload("TrashCategories").
			Find(&videos).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, pkg.ErrVideoNotFound
	}

	videoMap := make(map[int]video.Video)
	for _, v := range videos {
		videoMap[v.ID] = v
	}

	uniqueVideos := make([]video.Video, 0, len(videoMap))
	for _, v := range videoMap {
		uniqueVideos = append(uniqueVideos, v)
	}

	return &uniqueVideos, nil
}

func (repository *UserVideoRepositoryImpl) GetVideoDetail(id int) (*video.Video, *[]video.Comment, error) {
	var videos video.Video
	var comments []video.Comment

	if err := repository.DB.GetDB().
		Where("id = ?", id).
		Order("created_at desc").
		First(&videos).Error; err != nil {
		return nil, nil, err
	}

	if err := repository.DB.GetDB().
		Preload("User").
		Where("video_id = ?", id).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		return nil, nil, err
	}

	return &videos, &comments, nil
}

func (repository *UserVideoRepositoryImpl) AddComment(comment *video.Comment) error {
	if err := repository.DB.GetDB().Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (repository *UserVideoRepositoryImpl) UpdateViewer(view int, id int) error {
	if err := repository.DB.GetDB().Model(&video.Video{}).Where("id = ?", id).Update("viewer", view).Error; err != nil {
		return err
	}
	return nil
}
