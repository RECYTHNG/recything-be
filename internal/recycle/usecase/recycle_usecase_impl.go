package usecase

import (
	"github.com/sawalreverr/recything/internal/recycle/dto"
	"github.com/sawalreverr/recything/internal/recycle/repository"
)

type RecycleUsecaseImpl struct {
	RecycleRepository repository.RecycleRepository
}

func NewRecycleUsecaseImpl(repository repository.RecycleRepository) *RecycleUsecaseImpl {
	return &RecycleUsecaseImpl{
		RecycleRepository: repository,
	}
}

func (usecase *RecycleUsecaseImpl) GetHomeRecycleUsecase() (*dto.RecycleHomeResponse, error) {
	tasks, err := usecase.RecycleRepository.GetTasks()
	if err != nil {
		return nil, err
	}
	categories, err := usecase.RecycleRepository.GetCategoryVideos()
	if err != nil {
		return nil, err
	}
	videos, err := usecase.RecycleRepository.GetAllVideo()
	if err != nil {
		return nil, err
	}
	var dataTask []dto.DataTask
	var dataCategory []dto.DataCategory
	var dataVideo []dto.DataVideo
	data := &dto.RecycleHomeResponse{
		DataTask:     &dataTask,
		DataCategory: &dataCategory,
		DataVideo:    &dataVideo,
	}

	for _, task := range *tasks {
		dataTask = append(dataTask, dto.DataTask{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Thumbnail:   task.Thumbnail,
		})
	}
	for _, category := range *categories {
		dataCategory = append(dataCategory, dto.DataCategory{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	for _, video := range *videos {
		dataVideo = append(dataVideo, dto.DataVideo{
			Id:          video.ID,
			Title:       video.Title,
			Description: video.Description,
			Thumbnail:   video.Thumbnail,
			Link:        video.Link,
			Viewer:      video.Viewer,
		})
	}
	data.DataTask = &dataTask
	data.DataCategory = &dataCategory
	data.DataVideo = &dataVideo
	return data, nil
}

func (usecase *RecycleUsecaseImpl) SearchVideoUsecase(title string, category string) (*dto.SearchVideoResponse, error) {
	videos, err := usecase.RecycleRepository.SearchVideo(title, category)
	if err != nil {
		return nil, err
	}
	var dataVideo []dto.DataVideoSearch
	for _, video := range *videos {
		dataVideo = append(dataVideo, dto.DataVideoSearch{
			Id:          video.ID,
			Title:       video.Title,
			Description: video.Description,
			Thumbnail:   video.Thumbnail,
			Link:        video.Link,
			Viewer:      video.Viewer,
			Category:    dto.DataCategory{Id: video.VideoCategoryID, Name: video.Category.Name},
		})
	}
	data := &dto.SearchVideoResponse{
		DataVideo: &dataVideo,
	}
	return data, nil
}
