package pkg

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrStatusForbidden     = errors.New("forbidden")
	ErrStatusInternalError = errors.New("internal server error")
	ErrNoPrivilege         = errors.New("no permission to doing this task")

	// Authentication
	ErrEmailAlreadyExists       = errors.New("email already exists")
	ErrPhoneNumberAlreadyExists = errors.New("phone number already exists")
	ErrUserNotFound             = errors.New("user not found")
	ErrPasswordInvalid          = errors.New("password invalid")
	ErrOTPInvalid               = errors.New("otp invalid")
	ErrNeedToVerify             = errors.New("verify account false")
	ErrUserAlreadyVerified      = errors.New("user already verified")

	// Upload Cloudinary
	ErrUploadCloudinary       = errors.New("upload cloudinary server error")
	ErrUploadCloudinaryFailed = errors.New("failed to upload to cloudinary")
	UploadToCloudinary        = errors.New("upload to cloudinary")
	// admin
	ErrAdminNotFound = errors.New("admin not found")

	// article
	ErrTitleAlreadyExists = errors.New("title already exists")
	ErrArticleNotFound    = errors.New("article not found")

	//ID
	// GenerateID = = errors.New("failed to generate id")
)

func GenerateID() string {
	return uuid.New().String()
}
