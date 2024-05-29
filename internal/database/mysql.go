package database

import (
	"fmt"
	"log"
	"time"

	"github.com/sawalreverr/recything/config"
	adminEntity "github.com/sawalreverr/recything/internal/admin/entity"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/report"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDatabase struct {
	DB *gorm.DB
}

var (
	dbInstance *mysqlDatabase
)

func NewMySQLDatabase(conf *config.Config) Database {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	log.Println("Successfully connected to database:", conf.DB.DBName)

	dbInstance = &mysqlDatabase{DB: db}
	return dbInstance
}

func (m *mysqlDatabase) InitWasteMaterials() {
	initialWasteMaterials := []report.WasteMaterial{
		{ID: "MTR01", Type: "plastik"},
		{ID: "MTR02", Type: "kaca"},
		{ID: "MTR03", Type: "kayu"},
		{ID: "MTR04", Type: "kertas"},
		{ID: "MTR05", Type: "baterai"},
		{ID: "MTR06", Type: "besi"},
		{ID: "MTR07", Type: "limbah berbahaya"},
		{ID: "MTR08", Type: "limbah beracun"},
		{ID: "MTR09", Type: "sisa makanan"},
		{ID: "MTR10", Type: "tak terdeteksi"},
	}

	for _, material := range initialWasteMaterials {
		m.DB.FirstOrCreate(&material, material)
	}
}

func (m *mysqlDatabase) InitSuperAdmin() {
	hashed, _ := helper.GenerateHash("superadmin@123")

	admin := adminEntity.Admin{
		ID:        "AD0001",
		Name:      "John Doe Senior",
		Email:     "john.doe.sr@gmail.com",
		Password:  hashed,
		Role:      "super admin",
		ImageUrl:  "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.GetDB().Create(&admin)
	log.Println("Super admin data added!")
}

func (m *mysqlDatabase) GetDB() *gorm.DB {
	return dbInstance.DB
}
