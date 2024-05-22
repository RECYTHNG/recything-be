package dto

type Register struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10"`
	Password    string `json:"password" validate:"required,min=8"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
