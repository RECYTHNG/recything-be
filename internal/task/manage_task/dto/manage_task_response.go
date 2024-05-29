package dto

import "time"

type CreateTaskResponse struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Steps       []TaskSteps `json:"steps"`
}
