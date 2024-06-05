package dto

type CreateDataVideoRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description" validate:"required"`
	UrlThumbnail string `json:"url_thumbnail" validate:"required"`
	LinkVideo    string `json:"link_video" validate:"required"`
	CategoryId   int    `json:"category_id" validate:"required"`
}

type CreateCategoryVideoRequest struct {
	Name string `json:"name" validate:"required"`
}
