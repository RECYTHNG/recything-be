package entity

import "time"

type TrashCategoryCore struct {
	ID        string
	TrashType string
	CreatedAt time.Time
	UpdatedAt time.Time
}
