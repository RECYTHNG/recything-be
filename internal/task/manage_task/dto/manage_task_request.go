package dto

type CreateTaskResquest struct {
	Title        string      `json:"title" validate:"required"`
	Description  string      `json:"description" validate:"required"`
	ThumbnailUrl string      `json:"thumbnail_url" validate:"required"`
	StartDate    string      `json:"start_date" validate:"required"`
	EndDate      string      `json:"end_date" validate:"required"`
	Point        int         `json:"point" validate:"required"`
	TaskSteps    []TaskSteps `json:"task_steps" validate:"required"`
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
	StartDate    string      `json:"start_date" validate:"required"`
	EndDate      string      `json:"end_date" validate:"required"`
	Point        int         `json:"point" validate:"required"`
	TaskSteps    []TaskSteps `json:"task_steps" validate:"required"`
}
