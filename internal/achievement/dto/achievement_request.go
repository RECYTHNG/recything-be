package dto

type AchievementRequestCreate struct {
	Level       string `json:"level" validate:"required"`
	Lencana     string `json:"lencana" validate:"required"`
	TargetPoint int    `json:"target_point" validate:"required,min=1"`
}

type AchievementRequestUpdate struct {
	Level       string `json:"level" validate:"required"`
	Lencana     string `json:"lencana" validate:"required"`
	TargetPoint int    `json:"target_point" validate:"required,min=1"`
}
