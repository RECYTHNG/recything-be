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

	// Manage Task
	ErrTaskStepsNull = errors.New("steps cannot be null")
	ErrTaskNotFound  = errors.New("task not found")

	// User Task
	ErrImageTaskNull = errors.New("image task cannot be null")
	ErrUserTaskExist = errors.New("user task already exist")
)
