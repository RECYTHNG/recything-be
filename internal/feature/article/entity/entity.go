package entity

import "time"

type ArticleCore struct {
	ID          string
	Title       string
	Description string
	Categories  []ArticleTrashCategoryCore
	Category_id []string
	Thumbnail   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticleTrashCategoryCore struct {
	// TrashCategoryID string
	Category string
}
