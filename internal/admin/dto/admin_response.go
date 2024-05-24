package dto

type AdminResponseRegister struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	ProfilePhoto string `json:"profile_photo"`
}

type AdminDataGetAll struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type AdminResponseGetDataAll struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []AdminDataGetAll `json:"data"`
	Limit   int               `json:"limit"`
	Total   int               `json:"total"`
}
