package usecase

import (
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

func (usecase *AdminUsecaseImpl) AddAdminUsecase(request dto.AdminRequestCreate) (*entity.Admin, error) {
	findAdmin, _ := usecase.Repository.FindAdminByEmail(request.Email)
	if findAdmin != nil {
		return nil, pkg.ErrEmailAlreadyExist
	}

	hashPassword, _ := helper.GenerateHash(request.Password)

	admin, error := usecase.Repository.CreateDataAdmin(&entity.Admin{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashPassword,
		Role:      request.Role,
		DeletedAt: gorm.DeletedAt{},
	})
	if error != nil {
		return nil, error
	}
	return admin, nil
}

func (usecase *AdminUsecaseImpl) UpdateAdminUsecase(request dto.AdminUpdateRequest) (*entity.Admin, error) {
	if err := usecase.Validate.Struct(request); err != nil {
		return nil, err
	}
	admin, error := usecase.Repository.UpdateDataAdmin(&entity.Admin{
		ID:        "1",
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.NewPassword,
		Role:      request.Role,
		DeletedAt: gorm.DeletedAt{},
	}, request.OldPassword)
	if error != nil {
		return nil, error
	}
	return admin, nil
}

func (usecase *AdminUsecaseImpl) UploadProfileUsecase(request dto.UploadProfileImageRequest) (*entity.Admin, error) {
	_, err := usecase.Repository.FindAdminByID("1")

	if err != nil {
		return nil, err
	}

	admin, error := usecase.Repository.AddProfileImage("1", request.ImageUrl)
	if error != nil {
		return nil, error
	}
	return admin, nil
}
