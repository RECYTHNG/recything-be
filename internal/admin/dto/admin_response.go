package dto

type AdminResponseRegister struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UploadProfileImageResponse struct {
	ImageUrl string `json:"image_url"`
}
