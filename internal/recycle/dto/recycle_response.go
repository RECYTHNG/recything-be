package dto

type DataTask struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"url_thumbnail"`
}

type DataCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DataVideo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"url_thumbnail"`
	Link        string `json:"link_video"`
	Viewer      int    `json:"viewer"`
}

type RecycleHomeResponse struct {
	DataTask     *[]DataTask     `json:"data_task"`
	DataCategory *[]DataCategory `json:"data_category"`
	DataVideo    *[]DataVideo    `json:"data_video"`
}

type DataVideoSearch struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Thumbnail   string       `json:"url_thumbnail"`
	Link        string       `json:"link_video"`
	Viewer      int          `json:"viewer"`
	Category    DataCategory `json:"category"`
}

type SearchVideoResponse struct {
	DataVideo *[]DataVideoSearch `json:"data_video"`
}

type GetAllCategoryVideoResponse struct {
	DataCategory *[]DataCategory `json:"data_category"`
}
