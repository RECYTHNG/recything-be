package report

import (
	"time"

	"github.com/google/uuid"
	"github.com/sawalreverr/recything/internal/helper"
	rpt "github.com/sawalreverr/recything/internal/report"
	"github.com/sawalreverr/recything/pkg"
)

type reportUsecase struct {
	reportRepository rpt.ReportRepository
}

func NewReportUsecase(repo rpt.ReportRepository) rpt.ReportUsecase {
	return &reportUsecase{reportRepository: repo}
}

func (uc *reportUsecase) CreateReport(report rpt.ReportInput, authorID string, imageURLs []string) (*rpt.ReportDetail, error) {
	lastID, _ := uc.reportRepository.FindLastID()
	newID := helper.GenerateCustomID(lastID, "RPT")

	newReport := rpt.Report{
		ID:          newID,
		AuthorID:    authorID,
		ReportType:  report.ReportType,
		Title:       report.Title,
		Description: report.Description,
		WasteType:   report.WasteType,
		Latitude:    report.Latitude,
		Longitude:   report.Longitude,
		Address:     report.Address,
		City:        report.City,
		Province:    report.Province,
	}

	createdReport, err := uc.reportRepository.Create(newReport)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, materialType := range report.WasteMaterials {
		material, err := uc.reportRepository.FindWasteMaterialByType(materialType)
		if err != nil {
			_ = uc.reportRepository.Delete(createdReport.ID)
			return nil, err
		}

		reportMaterial := rpt.ReportWasteMaterial{
			ID:              uuid.New(),
			ReportID:        createdReport.ID,
			WasteMaterialID: material.ID,
		}

		if _, err := uc.reportRepository.AddReportMaterial(reportMaterial); err != nil {
			_ = uc.reportRepository.Delete(createdReport.ID)
			return nil, err
		}
	}

	for _, url := range imageURLs {
		reportImage := rpt.ReportImage{
			ID:       uuid.New(),
			ReportID: createdReport.ID,
			ImageURL: url,
		}

		if _, err := uc.reportRepository.AddImage(reportImage); err != nil {
			_ = uc.reportRepository.Delete(createdReport.ID)
			_ = uc.reportRepository.DeleteAllReportMaterial(createdReport.ID)
			return nil, err
		}
	}

	images, err := uc.reportRepository.FindAllImage(createdReport.ID)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	materials, err := uc.reportRepository.FindAllReportMaterial(createdReport.ID)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	reportDetail := rpt.ReportDetail{
		ID:             createdReport.ID,
		AuthorID:       createdReport.AuthorID,
		ReportType:     createdReport.ReportType,
		Title:          createdReport.Title,
		Description:    createdReport.Description,
		WasteType:      createdReport.WasteType,
		Latitude:       createdReport.Latitude,
		Longitude:      createdReport.Longitude,
		Address:        createdReport.Address,
		City:           createdReport.City,
		Province:       createdReport.Province,
		CreatedAt:      createdReport.CreatedAt,
		WasteMaterials: *materials,
		ReportImages:   *images,
	}

	return &reportDetail, nil
}

func (uc *reportUsecase) UpdateStatusReport(report rpt.UpdateStatus) error {
	reportFound, err := uc.reportRepository.FindByID(report.ID)
	if err != nil {
		return pkg.ErrReportNotFound
	}

	reportFound.Status = report.Status

	if err := uc.reportRepository.Update(*reportFound); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (uc *reportUsecase) FindAllReports(page, limit int, reportType, status string, date time.Time) (*[]rpt.ReportDetail, int64, error) {
	var reportDetails []rpt.ReportDetail
	reports, total, err := uc.reportRepository.FindAll(page, limit, reportType, status, date)
	if err != nil {
		return nil, 0, pkg.ErrStatusInternalError
	}

	for _, report := range *reports {
		images, err := uc.reportRepository.FindAllImage(report.ID)
		if err != nil {
			return nil, 0, pkg.ErrStatusInternalError
		}

		materials, err := uc.reportRepository.FindAllReportMaterial(report.ID)
		if err != nil {
			return nil, 0, pkg.ErrStatusInternalError
		}

		reportDetail := rpt.ReportDetail{
			ID:             report.ID,
			AuthorID:       report.AuthorID,
			ReportType:     report.ReportType,
			Title:          report.Title,
			Description:    report.Description,
			WasteType:      report.WasteType,
			Latitude:       report.Latitude,
			Longitude:      report.Longitude,
			Address:        report.Address,
			City:           report.City,
			Province:       report.Province,
			CreatedAt:      report.CreatedAt,
			WasteMaterials: *materials,
			ReportImages:   *images,
		}

		reportDetails = append(reportDetails, reportDetail)
	}

	return &reportDetails, total, nil
}
