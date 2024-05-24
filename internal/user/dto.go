package user

import "time"

type UserDetail struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	BirthDate   time.Time `json:"birth_date"`
	Address     string    `json:"address"`
}
