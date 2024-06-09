package dto

import "mime/multipart"

type UpdateAchievementRequest struct {
	Level       string                `json:"level"`
	TargetPoint int                   `json:"target_point"`
	Badge       *multipart.FileHeader `json:"-"`
}
