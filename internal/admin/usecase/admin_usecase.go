package usecase

import (
	"io"

	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/entity"
)

type AdminUsecase interface {
	AddAdminUsecase(request dto.AdminRequestCreate) (*entity.Admin, error)
	GetDataAllAdminUsecase(limit int, offset int) ([]entity.Admin, int, error)
	GetDataAdminByIdUsecase(id string) (*entity.Admin, error)
	UpdateAdminUsecase(request dto.AdminUpdateRequest, id string, file io.Reader) (*entity.Admin, error)
	DeleteAdminUsecase(id string) error
	GetDataAdminByEmailUsecase(email string) (*entity.Admin, error)
	GetProfileAdmin(id string) (*entity.Admin, error)
	UpdateAdminCurrenLoginUsecase(id string, request *dto.AdminUpdateRequest) (*entity.Admin, error)
}
