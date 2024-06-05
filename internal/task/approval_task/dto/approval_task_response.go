package dto

import "time"

type DataTasks struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type DataUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DataUserTask struct {
	Id            string    `json:"id"`
	StatusAccept  string    `json:"status_accept"`
	Point         int       `json:"point"`
	TaskChallenge DataTasks `json:"task"`
	User          DataUser  `json:"user"`
}

type GetUserTaskPagination struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      []DataUserTask `json:"data"`
	Page      int            `json:"page"`
	Limit     int            `json:"limit"`
	TotalData int            `json:"total_data"`
	TotalPage int            `json:"total_page"`
}

type GetUserTaskDetailsResponse struct {
	Id        string        `json:"id"`
	TitleTask string        `json:"title_task"`
	StartDate time.Time     `json:"start_date"`
	EndDate   time.Time     `json:"end_date"`
	UserName  string        `json:"user_name"`
	Images    []*DataImages `json:"images"`
}

type DataImages struct {
	Id          int       `json:"id"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	UploadedAt  time.Time `json:"uploaded_at"`
}
