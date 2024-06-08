package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type RecycleRepository interface {
	GetTasks() (*[]task.TaskChallenge, error)
	GetCategoryVideos() (*[]video.VideoCategory, error)
	GetAllVideo() (*[]video.Video, error)
	SearchVideo(title string, category string) (*[]video.Video, error)
	GetAllCategoryVideo() (*[]video.VideoCategory, error)
}
