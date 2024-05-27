package pkg

import "errors"

const (
	ERROR_DATA_ID          = "id tidak ditemukan"
	ERROR_ID_INVALID       = "id salah"
	ERROR_DATA_EMAIL       = "email tidak ditemukan"
	ERROR_FORMAT_EMAIL     = "error : format email tidak valid"
	ERROR_EMAIL_EXIST      = "error : email sudah digunakan"
	ERROR_AKSES_ROLE       = "akses ditolak"
	ERROR_PASSWORD         = "error : password lama tidak sesuai"
	ERROR_CONFIRM_PASSWORD = "error : konfirmasi password tidak sesuai"
	ERROR_ID_ROLE          = "id atau role tidak ditemukan"
	ERROR_GET_DATA         = "data tidak ditemukan"
	ERROR_EMPTY            = "error : harap lengkapi data dengan benar"
	ERROR_EMPTY_FILE       = "error : tidak ada file yang di upload"
	ERROR_DATA_NOT_FOUND   = "data tidak ditemukan"
	ERROR_DATA_EXIST       = "error : data sudah ada"
	ERROR_INVALID_ID       = "error: id tidak boleh sama"
	ERROR_INVALID_UPDATE   = "error: data harus berberbeda dengan data sebelumnya"
	ERROR_INVALID_INPUT    = "data yang diinput tidak sesuai"
	ERROR_NOT_FOUND        = "data tidak ditemukan"
)

const (
	BRONZE   = "bronze"
	PLATINUM = "platinum"
	SILVER   = "silver"
	GOLD     = "gold"
)

var (
	ErrStatusForbidden     = errors.New("forbidden")
	ErrStatusInternalError = errors.New("internal server error")
	ErrNoPrivilege         = errors.New("no permission to doing this task")
)
