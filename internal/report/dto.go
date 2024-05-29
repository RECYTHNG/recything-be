package report

import (
	"mime/multipart"
	"time"
)

type ReportInput struct {
	ReportType     string                  `json:"report_type" validate:"required,oneof='littering' 'rubbish'"`
	Title          string                  `json:"title" validate:"required,min=3,max=100"`
	Description    string                  `json:"description" validate:"required"`
	WasteType      string                  `json:"waste_type" validate:"required,oneof='sampah basah' 'sampah kering' 'sampah basah,sampah kering' 'organik' 'anorganik' 'berbahaya'"`
	WasteMaterials []string                `json:"waste_materials" form:"waste_materials" validate:"required,dive,oneof='plastik' 'kaca' 'kayu' 'kertas' 'baterai' 'besi' 'limbah berbahaya' 'limbah beracun' 'sisa makanan' 'tak terdeteksi'"`
	Latitude       float64                 `json:"latitude" validate:"required,latitude"`
	Longitude      float64                 `json:"longitude" validate:"required,longitude"`
	Address        string                  `json:"address" validate:"required"`
	City           string                  `json:"city" validate:"required"`
	Province       string                  `json:"province" validate:"required"`
	ReportImages   []*multipart.FileHeader `json:"-"`
}

type UpdateStatus struct {
	ID     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required,oneof='approved' 'reject'"`
}

type ReportDetail struct {
	ID          string  `json:"id"`
	AuthorID    string  `json:"author_id"`
	ReportType  string  `json:"report_type"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	WasteType   string  `json:"waste_type"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	Province    string  `json:"province"`
	Status      string  `json:"status"`

	CreatedAt      time.Time       `json:"created_at"`
	WasteMaterials []WasteMaterial `json:"waste_materials"`
	ReportImages   []string        `json:"report_images"`
}
