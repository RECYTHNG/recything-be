package dto

import "mime/multipart"

type UserTaskRequestCreate struct {
	TaskChallengeId string `json:"task_challenge_id" validate:"required"`
}

type UploadImageTask struct {
	Description string                  `json:"description" validate:"required"`
	Images      []*multipart.FileHeader `json:"-"`
}
