package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	a "github.com/sawalreverr/recything/internal/auth"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type authHandler struct {
	authUsecase a.AuthUsecase
}

func NewAuthHandler(uc a.AuthUsecase) a.AuthHandler {
	return &authHandler{authUsecase: uc}
}

func (h *authHandler) Register(c echo.Context) error {
	var request a.Register

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	newUser, err := h.authUsecase.RegisterUser(request)
	if err != nil {
		if errors.Is(err, pkg.ErrStatusInternalError) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusConflict, err.Error())
	}

	response := a.RegisterResponse{
		ID:         newUser.ID,
		Name:       newUser.Name,
		Email:      newUser.Email,
		IsVerified: newUser.IsVerified,
	}

	if err := helper.SendMail(newUser.Email, newUser.OTP); err != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "user successfully register! otp sent to your email", response)
}

func (h *authHandler) VerifyOTP(c echo.Context) error {
	var request a.OTPRequest

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.authUsecase.VerifyOTP(request); err != nil {
		if errors.Is(err, pkg.ErrStatusInternalError) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "otp successfully verified. registration complete!", nil)
}

func (h *authHandler) ResendOTP(c echo.Context) error {
	var request a.ResendOTP

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.authUsecase.UpdateOTP(request.Email); err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "new otp sent to your email!", nil)
}

func (h *authHandler) Login(c echo.Context) error {
	var request a.Login

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.authUsecase.LoginUser(request)
	if err != nil {
		if errors.Is(err, pkg.ErrStatusInternalError) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusUnauthorized, err.Error())
	}

	response := a.LoginResponse{
		Email: request.Email,
		Token: token,
	}

	return helper.ResponseHandler(c, http.StatusOK, "login successfully!", response)
}
