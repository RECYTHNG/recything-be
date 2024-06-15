package repository

type DashboardRepository interface {
	GetTotalUser() (int, int, error)
	GetTotalReport() (int, int, error)
	GetTotalChallenge() (int, int, error)
	GetTotalVideo() (int, int, error)
	GetTotalArticle() (int, int, error)
	GetUserClassic() (int, error)
	GetUserSilver() (int, error)
	GetUserGold() (int, error)
	GetUserPlatinum() (int, error)
	GetReportLittering() (int, error)
	GetReportRubbish() (int, error)
}
