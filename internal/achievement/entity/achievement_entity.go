package entity

import "time"

type Achievement struct {
	Id          int `gorm:"primaryKey"`
	Level       string
	Lencana     string
	TargetPoint int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
