package dto

type AdminRequestCreate struct {
	Name            string `json:"name" form:"name" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
	Role            string `json:"role" form:"role" validate:"required"`
	ProfileUrl      string `json:"profile_url" form:"profile_url" validate:"required"`
}

type AdminUpdateRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
	Role        string `json:"role" form:"role" validate:"required"`
	ProfileUrl  string `json:"profile_url" form:"profile_url" validate:"required"`
}
