package entity

import "time"

type ArticleCore struct {
	ID          string
	Title       string
	SubTitle    string
	Description string
	Image       string
	Categories  []ArticleTrashCategoryCore
	Category_id []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticleTrashCategoryCore struct {
	// TrashCategoryID string
	Category string
}
