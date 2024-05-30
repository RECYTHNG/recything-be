package dto

type ArticleResponseRegister struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type ArticleDataGetAll struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type ArticleResponseGetDataAll struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      []ArticleDataGetAll `json:"data"`
	Page      int                 `json:"page"`
	Limit     int                 `json:"limit"`
	TotalData int                 `json:"total_data"`
	TotalPage int                 `json:"total_page"`
}

type ArticleResponseGetDataById struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

type ArticleResponseUpdate struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
