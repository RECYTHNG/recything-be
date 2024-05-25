package usecase

import (
	"io"

	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/entity"
)

type AdminUsecase interface {
	AddAdminUsecase(request dto.AdminRequestCreate, file io.Reader) (*entity.Admin, error)
	GetDataAllAdminUsecase(limit int) ([]entity.Admin, int, error)
	GetDataAdminByIdUsecase(id string) (*entity.Admin, error)
	UpdateAdminUsecase(request dto.AdminUpdateRequest, id string, file io.Reader) (*entity.Admin, error)
	DeleteAdminUsecase(id string) error
}
