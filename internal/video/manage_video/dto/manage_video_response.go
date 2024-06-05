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
