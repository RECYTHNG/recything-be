package report

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	rpt "github.com/sawalreverr/recything/internal/report"
	"github.com/sawalreverr/recything/pkg"
)

type reportHandler struct {
	ReportUsecase rpt.ReportUsecase
}

func NewReportHandler(usecase rpt.ReportUsecase) rpt.ReportHandler {
	return &reportHandler{ReportUsecase: usecase}
}

func (h *reportHandler) NewReport(c echo.Context) error {
	var request rpt.ReportInput

	authorID := c.Get("user").(*helper.JwtCustomClaims).UserID

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	form, _ := c.MultipartForm()
	imageFiles := form.File["images"]

	validImages, err := helper.ImagesValidation(imageFiles)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	var imageURLs []string
	for _, file := range validImages {
		resultURL, err := helper.UploadToCloudinary(file, "recything/reports")
		if err != nil {
			helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}
		imageURLs = append(imageURLs, resultURL)
	}

	newReport, err := h.ReportUsecase.CreateReport(request, authorID, imageURLs)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusCreated, "report created!", newReport)
}

func (h *reportHandler) UpdateStatus(c echo.Context) error {
	var request rpt.UpdateStatus

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if err := h.ReportUsecase.UpdateStatusReport(request); err != nil {
		if errors.Is(err, pkg.ErrReportNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	return helper.ResponseHandler(c, http.StatusOK, "report status updated!", nil)
}

func (h *reportHandler) GetAllReports(c echo.Context) error {

	return helper.ResponseHandler(c, http.StatusOK, "ok", nil)
}
