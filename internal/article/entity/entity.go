package entity

import "time"

type ArticleCore struct {
	ID         string
	Title      string
	Categories []ArticleTrashCategory
	Thumbnail  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ArticleTrashCategory struct {
	// TrashCategoryID string
	Category string
}
