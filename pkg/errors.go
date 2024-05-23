package pkg

import "errors"

var (
	ErrStatusForbidden     = errors.New("forbidden")
	ErrStatusInternalError = errors.New("internal server error")
	ErrNoPrivilege         = errors.New("no permission to doing this task")
	ErrEmailAlreadyExist   = errors.New("email already exist")
	ErrFileTooLarge        = errors.New("upload image size must less than 2MB")
	ErrInvalidFileType     = errors.New("only image allowed")
	ErrUploadCloudinary    = errors.New("upload to cloudinary error")
	ErrDataNotFound        = errors.New("data not found")
)
