package model

import (
	"time"

	"gorm.io/gorm"
)

type TrashCategory struct {
	ID        string `gorm:"primaryKey"`
	TrashType string `gorm:not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *TrashCategory) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuidNew()
	t.ID = newUuid.String()
	return nil
}
