package dto

type CreateDataVideoRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	UrlThumbnail string `json:"url_thumbnail"`
	LinkVideo    string `json:"link_video"`
	Category     string `json:"category"`
}
