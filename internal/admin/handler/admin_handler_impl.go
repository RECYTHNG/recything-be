package handler

import (
	"errors"
	"net/http"
	"strconv"
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

	if request.Role != "admin" && request.Role != "super admin" {
		return helper.ErrorHandler(c, http.StatusBadRequest, "role must be admin or super admin")
	}

	admin, errUc := handler.Usecase.AddAdminUsecase(request)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrEmailAlreadyExists) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrEmailAlreadyExists.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+errUc.Error())
	}

	data := dto.AdminResponseRegister{
		Id:           admin.ID,
		Name:         admin.Name,
		Email:        admin.Email,
		Role:         admin.Role,
		ProfilePhoto: admin.ImageUrl,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusCreated, responseData)
}

func (handler *adminHandlerImpl) GetDataAllAdminHandler(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}

	limitInt, errLimit := strconv.Atoi(limit)
	if errLimit != nil || limitInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit parameter")
	}
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil || pageInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid page parameter")
	}

	admins, totalData, err := handler.Usecase.GetDataAllAdminUsecase(limitInt, pageInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	data := []dto.AdminDataGetAll{}
	for _, admin := range admins {
		data = append(data, dto.AdminDataGetAll{
			Id:    admin.ID,
			Name:  admin.Name,
			Email: admin.Email,
			Role:  admin.Role,
		})
	}

	totalPage := totalData / limitInt
	if totalData%limitInt != 0 {
		totalPage++
	}

	dataRes := dto.AdminResponseGetDataAll{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return c.JSON(http.StatusOK, dataRes)
}

func (handler *adminHandlerImpl) GetDataAdminByIdHandler(c echo.Context) error {
	id := c.Param("adminId")

	admin, err := handler.Usecase.GetDataAdminByIdUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrAdminNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	data := dto.AdminResponseGetDataById{
		Id:           admin.ID,
		Name:         admin.Name,
		Email:        admin.Email,
		Role:         admin.Role,
		ProfilePhoto: admin.ImageUrl,
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)

}

func (handler *adminHandlerImpl) UpdateAdminHandler(c echo.Context) error {
	id := c.Param("adminId")

	var request dto.AdminUpdateRequest
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if request.Role != "admin" && request.Role != "super admin" {
		return helper.ErrorHandler(c, http.StatusBadRequest, "role must be admin or super admin")
	}

	file, errFile := c.FormFile("profile_photo")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "profile_photo is required")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "file is too large")
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid file type")
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "failed to open file: "+errOpen.Error())
	}
	defer src.Close()

	admin, errUc := handler.Usecase.UpdateAdminUsecase(request, id, src)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrAdminNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrAdminNotFound.Error())
		}

		if errors.Is(errUc, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
		}

		if errors.Is(errUc, pkg.ErrPasswordInvalid) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrPasswordInvalid.Error())
		}

		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+errUc.Error())
	}

	data := dto.AdminResponseUpdate{
		Id:           id,
		Name:         admin.Name,
		Email:        admin.Email,
		Role:         admin.Role,
		ProfilePhoto: admin.ImageUrl,
	}
	responseData := helper.ResponseData(http.StatusOK, "data successfully updated", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *adminHandlerImpl) GetProfileAdminHandler(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)
	admin, err := handler.Usecase.GetProfileAdmin(claims.UserID)
	if err != nil {
		if errors.Is(err, pkg.ErrAdminNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	data := dto.AdminResponseGetDataById{
		Id:           admin.ID,
		Name:         admin.Name,
		Email:        admin.Email,
		Role:         admin.Role,
		ProfilePhoto: admin.ImageUrl,
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *adminHandlerImpl) DeleteAdminHandler(c echo.Context) error {
	id := c.Param("adminId")

	err := handler.Usecase.DeleteAdminUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrAdminNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, err.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}
	responseData := helper.ResponseData(http.StatusOK, "data successfully deleted", nil)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *adminHandlerImpl) UpdateAdminCurrentLoginHandler(c echo.Context) error {
	var request dto.AdminUpdateRequest
	claims := c.Get("user").(*helper.JwtCustomClaims)

	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if request.Role != "admin" && request.Role != "super admin" {
		return helper.ErrorHandler(c, http.StatusBadRequest, "role must be admin or super admin")
	}

	admin, errUc := handler.Usecase.UpdateAdminCurrenLoginUsecase(claims.UserID, &request)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrAdminNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrAdminNotFound.Error())
		}
		if errors.Is(errUc, pkg.ErrPasswordInvalid) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrPasswordInvalid.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+errUc.Error())
	}
	data := dto.AdminResponseUpdate{
		Id:           claims.UserID,
		Name:         admin.Name,
		Email:        admin.Email,
		Role:         admin.Role,
		ProfilePhoto: admin.ImageUrl,
	}
	responseData := helper.ResponseData(http.StatusOK, "data successfully updated", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *adminHandlerImpl) AddProfileAdminHandler(c echo.Context) error {
	file, errFile := c.FormFile("profile_photo")

	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "please upload your image!")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "file is too large")
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid file type")
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

	data := dto.AdminResponseAddProfile{
		ProfilePhoto: imageUrl,
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}
