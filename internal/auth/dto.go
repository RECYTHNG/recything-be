package auth

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

type OTPRequest struct {
	Email string `json:"phone_number" validate:"required,min=10"`
	OTP   uint   `json:"otp"`
}
