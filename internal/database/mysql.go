package database

import (
	"fmt"
	"log"
	"time"

	"github.com/sawalreverr/recything/config"
	adminEntity "github.com/sawalreverr/recything/internal/admin/entity"
	faqEntity "github.com/sawalreverr/recything/internal/faq"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/report"
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
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

	log.Println("Waste material data added!")
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

	m.GetDB().FirstOrCreate(&admin)
	log.Println("Super admin data added!")
}

func (m *mysqlDatabase) InitFaqs() {
	faqs := []faqEntity.FAQ{
		{ID: "FAQ01", Category: "profil", Question: "Bagaimana cara saya memperbarui informasi profil saya?", Answer: "Anda dapat memperbarui informasi profil Anda melalui menu 'Pengaturan Profil' di aplikasi. Klik ikon profil, pilih 'Pengaturan', dan edit informasi yang diperlukan."},
		{ID: "FAQ02", Category: "profil", Question: "Apakah saya bisa mengubah alamat email yang sudah terdaftar?", Answer: "Ya, Anda bisa mengubah alamat email Anda melalui menu 'Pengaturan Profil'. Namun, Anda mungkin perlu memverifikasi alamat email baru Anda."},
		{ID: "FAQ03", Category: "profil", Question: "Bagaimana cara mengganti foto profil saya?", Answer: "Untuk mengganti foto profil, buka 'Profil Saya', klik pada foto profil Anda saat ini, dan pilih foto baru dari galeri atau ambil foto baru dengan kamera."},

		{ID: "FAQ04", Category: "littering", Question: "Bagaimana cara melaporkan sampah yang tidak pada tempatnya?", Answer: "Anda dapat melaporkan sampah yang tidak pada tempatnya melalui fitur 'Laporkan Sampah' di aplikasi. Ambil foto sampah tersebut, tambahkan deskripsi singkat, dan kirim laporan Anda."},
		{ID: "FAQ05", Category: "littering", Question: "Apakah ada sanksi bagi yang membuang sampah sembarangan?", Answer: "Ya, sesuai dengan peraturan daerah, membuang sampah sembarangan dapat dikenakan denda atau sanksi lainnya. Silakan cek peraturan lokal untuk detailnya."},
		{ID: "FAQ06", Category: "littering", Question: "Apa yang terjadi setelah saya melaporkan sampah?", Answer: "Setelah Anda melaporkan sampah, tim kami akan memverifikasi laporan tersebut dan mengkoordinasikan pembersihan dengan pihak berwenang setempat."},

		{ID: "FAQ07", Category: "rubbish", Question: "Apa saja jenis-jenis sampah yang dapat didaur ulang?", Answer: "Jenis sampah yang dapat didaur ulang termasuk plastik, kertas, kaca, dan logam. Pastikan untuk memisahkan sampah sesuai kategori sebelum mendaur ulang."},
		{ID: "FAQ08", Category: "rubbish", Question: "Bagaimana cara memisahkan sampah dengan benar?", Answer: "Pisahkan sampah berdasarkan jenisnya - organik, anorganik, dan berbahaya. Gunakan tempat sampah yang berbeda untuk setiap kategori untuk mempermudah proses daur ulang."},
		{ID: "FAQ09", Category: "rubbish", Question: "Apa yang dimaksud dengan sampah organik?", Answer: "Sampah organik adalah sampah yang berasal dari bahan-bahan alami yang dapat terurai, seperti sisa makanan, daun, dan potongan kayu."},

		{ID: "FAQ10", Category: "misi", Question: "Bagaimana cara berpartisipasi dalam misi kebersihan?", Answer: "Anda dapat berpartisipasi dalam misi kebersihan dengan mendaftar melalui aplikasi di bagian 'Misi'. Pilih misi yang tersedia dan ikuti instruksi yang diberikan."},
		{ID: "FAQ11", Category: "misi", Question: "Apa saja manfaat mengikuti misi kebersihan?", Answer: "Manfaat mengikuti misi kebersihan termasuk mendapatkan poin dan level, membantu menjaga lingkungan, dan berkesempatan memenangkan penghargaan."},
		{ID: "FAQ12", Category: "misi", Question: "Bagaimana cara menyelesaikan misi dan mendapatkan poin?", Answer: "Untuk menyelesaikan misi, ikuti semua instruksi yang diberikan dan laporkan hasil kerja Anda melalui aplikasi. Poin akan diberikan berdasarkan kontribusi Anda."},

		{ID: "FAQ13", Category: "lokasi sampah", Question: "Bagaimana cara menemukan tempat sampah terdekat?", Answer: "Anda dapat menemukan tempat sampah terdekat menggunakan fitur 'Cari Tempat Sampah' di aplikasi. Aplikasi akan menunjukkan lokasi tempat sampah di peta."},
		{ID: "FAQ14", Category: "lokasi sampah", Question: "Apa yang harus saya lakukan jika tidak menemukan tempat sampah di sekitar saya?", Answer: "Jika Anda tidak menemukan tempat sampah di sekitar Anda, simpan sampah Anda sampai Anda menemukan tempat yang sesuai untuk membuangnya atau laporkan kebutuhan tempat sampah baru melalui aplikasi."},
		{ID: "FAQ15", Category: "lokasi sampah", Question: "Apakah lokasi tempat sampah di aplikasi selalu diperbarui?", Answer: "Ya, kami berusaha untuk selalu memperbarui lokasi tempat sampah di aplikasi berdasarkan laporan pengguna dan data dari pihak berwenang setempat."},

		{ID: "FAQ16", Category: "poin dan level", Question: "Bagaimana cara mendapatkan poin?", Answer: "Anda bisa mendapatkan poin dengan menyelesaikan misi, melaporkan sampah, dan berpartisipasi dalam kegiatan kebersihan. Poin akan otomatis ditambahkan ke akun Anda."},
		{ID: "FAQ17", Category: "poin dan level", Question: "Apa yang bisa saya lakukan dengan poin yang saya kumpulkan?", Answer: "Poin yang Anda kumpulkan bisa ditukar dengan berbagai hadiah, diskon, atau digunakan untuk meningkatkan level akun Anda dalam aplikasi."},
		{ID: "FAQ18", Category: "poin dan level", Question: "Bagaimana cara meningkatkan level saya?", Answer: "Tingkatkan level Anda dengan mengumpulkan poin dari berbagai aktivitas dalam aplikasi. Setiap level baru memberikan akses ke fitur dan penghargaan tambahan."},

		{ID: "FAQ19", Category: "artikel", Question: "Di mana saya bisa membaca artikel terkait daur ulang dan kebersihan?", Answer: "Anda bisa membaca artikel terkait daur ulang dan kebersihan di bagian 'Artikel' dalam aplikasi. Kami menyediakan berbagai artikel informatif untuk membantu Anda lebih peduli terhadap lingkungan."},
		{ID: "FAQ20", Category: "artikel", Question: "Apakah artikel di aplikasi diperbarui secara berkala?", Answer: "Ya, artikel di aplikasi diperbarui secara berkala dengan konten terbaru mengenai daur ulang, tips kebersihan, dan informasi lingkungan lainnya."},
		{ID: "FAQ21", Category: "artikel", Question: "Bisakah saya berkontribusi menulis artikel untuk aplikasi?", Answer: "Tentu saja! Kami menerima kontribusi dari pengguna. Jika Anda tertarik, silakan hubungi kami melalui fitur 'Kontak Kami' di aplikasi untuk informasi lebih lanjut tentang cara berkontribusi."},
	}

	for _, faq := range faqs {
		m.GetDB().FirstOrCreate(&faq, faq)
	}
	log.Println("FAQs data added!")
}

func (m *mysqlDatabase) InitVideoCategories() {
	videoCategories := []video.VideoCategory{
		{Name: "Tips"},
		{Name: "Daur Ulang"},
		{Name: "Tutorial"},
		{Name: "Edukasi"},
		{Name: "Kampanye"},
		{Name: "Lainnya"},
	}
	for _, videoCategory := range videoCategories {
		m.GetDB().FirstOrCreate(&videoCategory, videoCategory)
	}
	log.Println("Video categories data added!")
}

func (m *mysqlDatabase) GetDB() *gorm.DB {
	return dbInstance.DB
}
