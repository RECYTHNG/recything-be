package dto

import "time"

type CreateTaskResquest struct {
	Title        string      `json:"title" validate:"required"`
	Description  string      `json:"description" validate:"required"`
	ThumbnailUrl string      `json:"thumbnail_url" validate:"required"`
	StartDate    time.Time   `json:"start_date" validate:"required"`
	EndDate      time.Time   `json:"end_date" validate:"required"`
	Steps        []TaskSteps `json:"steps" validate:"required"`
}

type TaskSteps struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateTaskRequest struct {
	Title        string      `json:"title" validate:"required"`
	Description  string      `json:"description" validate:"required"`
	ThumbnailUrl string      `json:"thumbnail_url" validate:"required"`
	StartDate    time.Time   `json:"start_date" validate:"required"`
	EndDate      time.Time   `json:"end_date" validate:"required"`
	Steps        []TaskSteps `json:"steps" validate:"required"`
}
