package repository

import (
	"github.com/sawalreverr/recything/internal/database"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type RecycleRepositoryImpl struct {
	Database database.Database
}

func NewRecycleRepositoryImpl(database database.Database) *RecycleRepositoryImpl {
	return &RecycleRepositoryImpl{
		Database: database,
	}
}

func (repository *RecycleRepositoryImpl) GetTasks() (*[]task.TaskChallenge, error) {
	var tasks []task.TaskChallenge
	if err := repository.Database.GetDB().
		Limit(2).
		Order("id desc").
		Find(&tasks).
		Error; err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (repository *RecycleRepositoryImpl) GetCategoryVideos() (*[]video.VideoCategory, error) {
	var categories []video.VideoCategory
	if err := repository.Database.GetDB().
		Limit(4).
		Order("id desc").
		Find(&categories).
		Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

func (repository *RecycleRepositoryImpl) GetAllVideo() (*[]video.Video, error) {
	var videos []video.Video
	if err := repository.Database.GetDB().
		Limit(2).
		Order("id desc").
		Find(&videos).
		Error; err != nil {
		return nil, err
	}

	return &videos, nil
}
