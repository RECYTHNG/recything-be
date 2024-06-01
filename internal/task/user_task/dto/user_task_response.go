package dto

import (
	"time"
)

type UserTaskGetAllResponse struct {
	Data []DataUserTask `json:"data"`
}

type DataUserTask struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Point       int         `json:"point"`
	Status      bool        `json:"status"`
	TaskSteps   []TaskSteps `json:"task_steps"`
}

type TaskSteps struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
