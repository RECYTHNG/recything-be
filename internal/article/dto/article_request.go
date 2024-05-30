package dto

type ArticleRequestCreate struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type ArticleUpdateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
