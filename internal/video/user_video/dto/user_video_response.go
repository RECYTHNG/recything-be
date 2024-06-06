package dto

import "time"

type DataVideo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	UrlThumbnail string `json:"url_thumbnail"`
	LinkVideo    string `json:"link_video"`
	Viewer       int    `json:"viewer"`
}

type GetAllVideoResponse struct {
	DataVideo []DataVideo `json:"data_video"`
}

type DataComment struct {
	Id        int       `json:"id"`
	Comment   string    `json:"comment"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetDetailsDataVideoByIdResponse struct {
	DataVideo *DataVideo     `json:"data_video"`
	Comments  *[]DataComment `json:"comments"`
}
