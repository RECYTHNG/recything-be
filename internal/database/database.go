package database

import "gorm.io/gorm"

type Database interface {
	GetDB() *gorm.DB
	InitSuperAdmin()
	InitWasteMaterials()
	InitFaqs()
	InitCustomDatas()
}
