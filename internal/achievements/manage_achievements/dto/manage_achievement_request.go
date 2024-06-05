package dto

type CreateArchievementRequest struct {
	Level       string `json:"level" validate:"required"`
	TargetPoint int    `json:"target_point" validate:"required"`
	BadgeUrl    string `json:"badge_url" validate:"required"`
}

type UpdateAchievementRequest struct {
	Level       string `json:"level" validate:"required"`
	TargetPoint int    `json:"target_point" validate:"required"`
	BadgeUrl    string `json:"badge_url" validate:"required"`
}
