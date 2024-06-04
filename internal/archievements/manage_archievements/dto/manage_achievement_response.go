package dto

type CreateArchievementResponse struct {
	Level       string `json:"level"`
	TargetPoint int    `json:"target_point"`
	BadgeUrl    string `json:"badge_url"`
}

type UploadBadgeResponse struct {
	BadgeUrl string `json:"badge_url"`
}
