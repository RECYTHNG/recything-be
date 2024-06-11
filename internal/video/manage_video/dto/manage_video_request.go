package dto

import "mime/multipart"

type CreateDataVideoRequest struct {
	Title           string                `json:"title" validate:"required"`
	Description     string                `json:"description" validate:"required"`
	LinkVideo       string                `json:"link_video" validate:"required"`
	VideoCategories []DataCategoryVideo   `json:"video_categories"`
	TrashCategories []DataTrashCategory   `json:"trash_categories"`
	Thumbnail       *multipart.FileHeader `json:"-"`
}

type DataCategoryVideo struct {
	Name string `json:"name"`
}

type DataTrashCategory struct {
	Name string `json:"name"`
}

type CreateCategoryVideoRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateDataVideoRequest struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	LinkVideo   string                `json:"link_video"`
	CategoryId  int                   `json:"category_id"`
	Thumbnail   *multipart.FileHeader `json:"-"`
}
