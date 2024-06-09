package dto

import "mime/multipart"

type UpdateAchievementRequest struct {
	Level       string                `json:"level" validate:"required"`
	TargetPoint int                   `json:"target_point" validate:"required"`
	Badge       *multipart.FileHeader `json:"-"`
}
