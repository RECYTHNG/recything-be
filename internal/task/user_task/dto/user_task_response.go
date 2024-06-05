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

type UserTaskResponseCreate struct {
	Id             string                      `json:"id"`
	StatusProgress string                      `json:"status_progress"`
	TaskChalenge   TaskChallengeResponseCreate `json:"task_challenge"`
}

type TaskChallengeResponseCreate struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Point       int         `json:"point"`
	StatusTask  bool        `json:"status_task"`
	TaskSteps   []TaskSteps `json:"task_steps"`
}

type UserTaskUploadImageResponse struct {
	Id             string             `json:"id"`
	StatusProgress string             `json:"status_progress"`
	StatusAccept   string             `json:"status_accepted"`
	Point          int                `json:"point"`
	TaskChallenge  DataTaskChallenges `json:"task_challenge"`
	Images         []Images           `json:"images"`
}

type Images struct {
	Images string `json:"images"`
}

type UserTaskGetByIdUserResponse struct {
	Id             string                      `json:"id"`
	StatusProgress string                      `json:"status_progress"`
	TaskChallenge  TaskChallengeResponseCreate `json:"task_challenge"`
}

type GetUserTaskDoneByIdUserResponse struct {
	Id             string             `json:"id"`
	StatusProgress string             `json:"status_progress"`
	StatusAccept   string             `json:"status_accepted"`
	Point          int                `json:"point"`
	ReasonReject   string             `json:"reason_reject"`
	TaskChallenge  DataTaskChallenges `json:"task_challenge"`
}

type DataTaskChallenges struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	StatusTask  bool        `json:"status_task"`
	TaskSteps   []TaskSteps `json:"task_steps"`
}

type GetUserTaskDetailsResponse struct {
	Id        string        `json:"id"`
	TitleTask string        `json:"title_task"`
	UserName  string        `json:"user_name"`
	Images    []*DataImages `json:"images"`
}

type DataImages struct {
	Id          int    `json:"id"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}