package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type adminHandlerImpl struct {
	Usecase usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) *adminHandlerImpl {
	return &adminHandlerImpl{Usecase: adminUsecase}
}

func (handler *adminHandlerImpl) AddAdminHandler(c echo.Context) error {
	var request dto.AdminRequestCreate
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	admin, errUc := handler.Usecase.AddAdminUsecase(request)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrEmailAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrEmailAlreadyExist.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	data := dto.AdminResponseRegister{
		Id:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Role:  admin.Role,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusCreated, responseData)

}

func (handler *adminHandlerImpl) UploadProfileHandler(c echo.Context) error {
	const maxFileSize = 2 * 1024 * 1024

	if !strings.HasPrefix(c.Request().Header.Get("Content-Type"), "multipart/form-data") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "request Content-Type isn't multipart/form-data")
	}

	file, errFile := c.FormFile("image_url")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "failed to get form file: "+errFile.Error())
	}

	if file.Size > maxFileSize {
		return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrFileTooLarge.Error())
	}

	fileType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(fileType, "image/") {
		return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrInvalidFileType.Error())
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "failed to open file: "+errOpen.Error())
	}
	defer src.Close()

	imageUrl, errUpload := helper.UploadToCloudinary(src, "profile_admin")
	if errUpload != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
	}

	// Memperbarui profil admin dengan URL gambar
	admin, errUploadProfile := handler.Usecase.UploadProfileUsecase(dto.UploadProfileImageRequest{ImageUrl: imageUrl})
	if errUploadProfile != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, errUploadProfile.Error())
	}

	// Membuat respons
	data := dto.UploadProfileImageResponse{
		ImageUrl: admin.ImageUrl,
	}

	// Mengirimkan respons
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", data))
}

func (handler *adminHandlerImpl) UpdateAdminHandler(c echo.Context) error {
	return nil
}
