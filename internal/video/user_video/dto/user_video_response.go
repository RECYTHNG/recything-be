package dto

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
