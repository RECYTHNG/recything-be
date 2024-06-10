package article

import (
	"time"
)

type ArticleInput struct {
	Title        string                `json:"title"`
	Description  string                `json:"description"`
	ThumbnailURL string                `json:"thumbnail_url"`
	Categories   []string              `json:"categories"`
	Sections     []ArticleSectionInput `json:"sections"`
}

type ArticleSectionInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type AdminDetail struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ArticleDetail struct {
	ID           string      `json:"id"`
	AuthorID     AdminDetail `json:"author"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	ThumbnailURL string      `json:"thumbnail_url"`
	CreatedAt    time.Time   `json:"created_at"`

	WasteCategories   []WasteCategory   `json:"waste_categories"`
	ContentCategories []ContentCategory `json:"content_categories"`
	Sections          []ArticleSection  `json:"sections"`
}

type ArticleResponsePagination struct {
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	Limit    int             `json:"limit"`
	Articles []ArticleDetail `json:"articles"`
}
