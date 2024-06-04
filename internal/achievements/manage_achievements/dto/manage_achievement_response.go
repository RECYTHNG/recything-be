package dto

type CreateArchievementResponse struct {
	Level       string `json:"level"`
	TargetPoint int    `json:"target_point"`
	BadgeUrl    string `json:"badge_url"`
}

type UploadBadgeResponse struct {
	BadgeUrl string `json:"badge_url"`
}

type DataAchievement struct {
	Id          int    `json:"id"`
	Level       string `json:"level"`
	TargetPoint int    `json:"target_point"`
	BadgeUrl    string `json:"badge_url"`
}

type GetAllAchievementResponse struct {
	Data []*DataAchievement `json:"data"`
}
