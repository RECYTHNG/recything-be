package report

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// struct
type Report struct {
	ID          string  `json:"id" gorm:"primaryKey"`
	AuthorID    string  `json:"author_id"`
	ReportType  string  `json:"report_type" gorm:"type:enum('littering', 'rubbish');"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	WasteType   string  `json:"waste_type" gorm:"type:enum('sampah basah', 'sampah kering', 'organik', 'anorganik', 'berbahaya');" `
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Address     string  `json:"address"`
	Location    string  `json:"location"`
	Status      string  `json:"status" gorm:"type:enum('need review', 'approve', 'reject');default:'need review'"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type WasteMaterial struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Type string `json:"type" gorm:"type:enum('plastik', 'kaca', 'kayu', 'kertas', 'baterai', 'besi', 'limbah berbahaya', 'limbah beracun', 'sisa makanan', 'tak terdeteksi');" `

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ReportWasteMaterial struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	ReportID        string    `json:"report_id"`
	WasteMaterialID string    `json:"waste_material_id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ReportImage struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	ReportID string    `json:"report_id"`
	ImageURL string    `json:"image_url"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// interface
type ReportRepository interface {
	Create(report Report) (*Report, error)
	FindByID(reportID string) (*Report, error)
	FindAll() (*[]Report, error)
	FindLastID() (string, error)
	Update(report Report) error
	Delete(reportID string) error

	AddImage(image ReportImage) (*ReportImage, error)
	DeleteImage(imageID string, reportID string) error
	DeleteAllImage(reportID string) error
	FindAllImage(reportID string) (*[]ReportImage, error)

	AddReportMaterial(material ReportWasteMaterial) (*ReportWasteMaterial, error)
	DeleteAllReportMaterial(reportID string) error
	FindAllReportMaterial(reportID string) (*[]ReportWasteMaterial, error)
}

type ReportUsecase interface {
}

type ReportHandler interface {
}
