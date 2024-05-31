package dto

import "time"

type AchievementResponse struct {
	Id          int       `json:"id"`
	Level       string    `json:"level"`
	Lencana     string    `json:"lencana"`
	TargetPoint int       `json:"target_point"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AchievementResponseGetAll struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      []AchievementResponse `json:"data"`
	Page      int                   `json:"page"`
	Limit     int                   `json:"limit"`
	TotalData int                   `json:"total_data"`
	TotalPage int                   `json:"total_page"`
}
