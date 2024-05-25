package usecase

import (
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/admin/repository"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type AdminUsecaseImpl struct {
	Repository repository.AdminRepository
	Validate   *validator.Validate
}

func NewAdminUsecase(adminRepo repository.AdminRepository) *AdminUsecaseImpl {
	return &AdminUsecaseImpl{Repository: adminRepo}
}

func (usecase *AdminUsecaseImpl) AddAdminUsecase(request dto.AdminRequestCreate, file io.Reader) (*entity.Admin, error) {
	findAdmin, _ := usecase.Repository.FindAdminByEmail(request.Email)
	if findAdmin != nil {
		return nil, pkg.ErrEmailAlreadyExists
	}

	imageUrl, errUpload := helper.UploadToCloudinary(file, "profile_admin")
	if errUpload != nil {
		return nil, pkg.ErrUploadCloudinary
	}

	findLastId, _ := usecase.Repository.FindLastIdAdmin()
	id := helper.GenerateCustomID(findLastId, "AD")

	hashPassword, _ := helper.GenerateHash(request.Password)

	admin := &entity.Admin{
		ID:        id,
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashPassword,
		Role:      request.Role,
		ImageUrl:  imageUrl,
		DeletedAt: gorm.DeletedAt{},
	}

	if _, err := usecase.Repository.CreateDataAdmin(admin); err != nil {
		return nil, err
	}
	return admin, nil
}

func (usecase *AdminUsecaseImpl) GetDataAllAdminUsecase(limit int) ([]entity.Admin, error) {
	admins, err := usecase.Repository.GetDataAllAdmin(limit)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (usecase *AdminUsecaseImpl) GetDataAdminByIdUsecase(id string) (*entity.Admin, error) {
	admin, err := usecase.Repository.FindAdminByID(id)

	if err != nil {
		return nil, pkg.ErrAdminNotFound
	}
	return admin, nil
}

func (usecase *AdminUsecaseImpl) UpdateAdminUsecase(request dto.AdminUpdateRequest, id string, file io.Reader) (*entity.Admin, error) {
	findAdmin, _ := usecase.Repository.FindAdminByID(id)
	if findAdmin == nil {
		return nil, pkg.ErrAdminNotFound
	}

	if matchPassword := helper.ComparePassword(findAdmin.Password, request.OldPassword); matchPassword == false {
		return nil, pkg.ErrPasswordInvalid
	}

	hashPassword, _ := helper.GenerateHash(request.NewPassword)

	imageUrl, errUpload := helper.UploadToCloudinary(file, "profile_admin_update")
	if errUpload != nil {
		return nil, pkg.ErrUploadCloudinary
	}

	admin, error := usecase.Repository.UpdateDataAdmin(&entity.Admin{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashPassword,
		Role:      request.Role,
		ImageUrl:  imageUrl,
		DeletedAt: gorm.DeletedAt{},
	}, id)
	if error != nil {
		return nil, error
	}
	return admin, nil
}
