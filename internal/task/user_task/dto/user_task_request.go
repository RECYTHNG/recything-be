package dto

import "mime/multipart"

type UploadImageTask struct {
	Description string                  `json:"description" validate:"required"`
	Images      []*multipart.FileHeader `json:"-"`
}

type UpdateUserTaskRequest struct {
	Description string                  `json:"description" validate:"required"`
	Images      []*multipart.FileHeader `json:"-"`
}
