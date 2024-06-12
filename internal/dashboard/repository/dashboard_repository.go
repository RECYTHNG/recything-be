package repository

type DashboardRepository interface {
	GetTotalUser() (int, error)
	GetTotalReport() (int, error)
	GetTotalChange() (int, error)
	GetTotalContent() (int, error)
	GetUserClassic() (int, error)
	GetUserSilver() (int, error)
	GetUserGold() (int, error)
	GetUserPlatinum() (int, error)
	GetReportLittering() (int, error)
	GetReportRubbish() (int, error)
}
