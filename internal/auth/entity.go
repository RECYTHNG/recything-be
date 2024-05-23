package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/user"
)

type AuthUsecase interface {
	RegisterUser(user Register) (*user.User, error)
	LoginUser(user Login) (string, error)
	VerifyOTP(user OTPRequest) error
	UpdateOTP(email string) error
}

type AuthHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	VerifyOTP(c echo.Context) error
	ResendOTP(c echo.Context) error
}
