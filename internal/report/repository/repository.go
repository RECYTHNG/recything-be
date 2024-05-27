package report

import (
	"github.com/sawalreverr/recything/internal/database"
	rpt "github.com/sawalreverr/recything/internal/report"
)

type reportRepository struct {
	DB database.Database
}

func NewReportRepository(db database.Database) rpt.ReportRepository {
	return &reportRepository{DB: db}
}

// Report
func (r *reportRepository) Create(report rpt.Report) (*rpt.Report, error) {
	if err := r.DB.GetDB().Create(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *reportRepository) FindByID(reportID string) (*rpt.Report, error) {
	var report rpt.Report
	if err := r.DB.GetDB().Where("id = ?", reportID).First(&report).Error; err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *reportRepository) FindLastID() (string, error) {
	var report rpt.Report
	if err := r.DB.GetDB().Order("id DESC").First(&report).Error; err != nil {
		return "RPT0000", err
	}

	return report.ID, nil
}

func (r *reportRepository) Update(report rpt.Report) error {
	if err := r.DB.GetDB().Save(&report).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) Delete(reportID string) error {
	var report rpt.Report
	if err := r.DB.GetDB().Where("id = ?", reportID).Delete(&report).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) FindAll() (*[]rpt.Report, error) {

	return nil, nil
}

// Report Image
func (r *reportRepository) AddImage(image rpt.ReportImage) (*rpt.ReportImage, error) {
	if err := r.DB.GetDB().Create(&image).Error; err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *reportRepository) DeleteImage(imageID string, reportID string) error {
	var reportImage rpt.ReportImage
	if err := r.DB.GetDB().Where("id = ? AND report_id = ?", imageID, reportID).Delete(&reportImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) DeleteAllImage(reportID string) error {
	var reportImage rpt.ReportImage
	if err := r.DB.GetDB().Where("report_id = ?", reportID).Delete(&reportImage).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) FindAllImage(reportID string) (*[]rpt.ReportImage, error) {
	var reportImages []rpt.ReportImage
	if err := r.DB.GetDB().Where("report_id = ?", reportID).Find(&reportImages).Error; err != nil {
		return nil, err
	}

	return &reportImages, nil
}

// Report Waste Materials
func (r *reportRepository) AddReportMaterial(material rpt.ReportWasteMaterial) (*rpt.ReportWasteMaterial, error) {
	if err := r.DB.GetDB().Create(&material).Error; err != nil {
		return nil, err
	}

	return &material, nil
}

func (r *reportRepository) DeleteAllReportMaterial(reportID string) error {
	var reportMaterial rpt.ReportWasteMaterial
	if err := r.DB.GetDB().Where("report_id = ?", reportID).Delete(&reportMaterial).Error; err != nil {
		return err
	}

	return nil
}

func (r *reportRepository) FindAllReportMaterial(reportID string) (*[]rpt.ReportWasteMaterial, error) {
	var reportMaterials []rpt.ReportWasteMaterial
	if err := r.DB.GetDB().Where("report_id = ?", reportID).Find(&reportMaterials).Error; err != nil {
		return nil, err
	}

	return &reportMaterials, nil
}
