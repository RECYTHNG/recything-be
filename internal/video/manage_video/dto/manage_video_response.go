package dto

type GetAllCategoryVideoResponse struct {
	Data []*DataCategory
}

type DataCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UploadThumbnailResponse struct {
	UrlThumbnail string `json:"url_thumbnail"`
}

type DataVideo struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	UrlThumbnail string `json:"url_thumbnail"`
}

type GetAllDataVideoPaginationResponse struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	Data      []*DataVideo `json:"data"`
	Page      int          `json:"page"`
	Limit     int          `json:"limit"`
	TotalData int          `json:"total_data"`
	TotalPage int          `json:"total_page"`
}

type GetDetailsDataVideoByIdResponse struct {
	Id           int          `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	UrlThumbnail string       `json:"url_thumbnail"`
	LinkVideo    string       `json:"link_video"`
	Viewer       int          `json:"viewer"`
	Category     DataCategory `json:"category"`
}
