package aboutus

type AboutUsResponse struct {
	ID          string                 `json:"id" gorm:"primaryKey"`
	Category    string                 `json:"category"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Images      []AboutUsImageResponse `json:"images"`
}

type AboutUsImageResponse struct {
	AboutUsID string `json:"about_us_id"`
	ImageURL  string `json:"image_url"`
}
