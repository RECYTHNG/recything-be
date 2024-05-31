package pkg

import "errors"

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
	ErrUploadCloudinary = errors.New("upload cloudinary server error")

	// admin
	ErrAdminNotFound = errors.New("admin not found")

	// article
	ErrTitleAlreadyExists = errors.New("title already exists")
	ErrArticleNotFound    = errors.New("article not found")

	// Achievement
	ErrAchievementNotFound = errors.New("achievement not found")
	ErrInvalidID           = errors.New("invalid id")
)
