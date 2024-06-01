package dto

type UserTaskRequestCreate struct {
	TaskChallengeId string `json:"task_challenge_id" validate:"required"`
}

type UserImageTask struct {
	ImageUrl string `json:"image_url" validate:"required"`
}
