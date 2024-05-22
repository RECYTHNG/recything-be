package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Email       string
	PhoneNumber string
	Password    string `json:"-"`
	Point       uint
	Gender      string    `gorm:"type:enum('laki-laki', 'perempuan')"`
	BirthDate   time.Time `gorm:"type:datetime"`
	Address     string    `json:"address"`
	PictureURL  string    `json:"picture_url"`
	OTP         uint      `json:"otp"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
