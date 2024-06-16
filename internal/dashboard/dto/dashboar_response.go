package dto

type DashboardResponse struct {
	User                TotalUser           `json:"user"`
	Report              TotalReport         `json:"report"`
	Challenge           TotalChallenge      `json:"challenge"`
	Content             TotalContent        `json:"content"`
	UserAchievement     UserAchievement     `json:"user_achievement"`
	TotalLittering      int                 `json:"total_littering"`
	TotalRubbish        int                 `json:"total_rubbish"`
	DataUserByAddress   []DataUserByAddress `json:"data_user_by_address"`
	DataReportLittering ReportLittering     `json:"data_report_littering"`
	DataReportRubbish   ReportRubbish       `json:"data_report_rubbish"`
}

type TotalUser struct {
	TotalUser                  int `json:"total_user"`
	AdditionUserSinceYesterday int `json:"addition_user_since_yesterday"`
}

type UserAchievement struct {
	TotalUser int `json:"total_user"`
	Classic   int `json:"classic"`
	Silver    int `json:"silver"`
	Gold      int `json:"gold"`
	Platinum  int `json:"platinum"`
}

type TotalReport struct {
	TotalReport                  int `json:"total_report"`
	AdditionReportSinceYesterday int `json:"addition_report_since_yesterday"`
}

type TotalChallenge struct {
	TotalChallenge                 int `json:"total_challenge"`
	AdditionChallengeSinceLastWeek int `json:"addition_challenge_since_last_week"`
}

type TotalContent struct {
	TotalContent         int `json:"total_content"`
	AdditionContentToday int `json:"addition_content_today"`
}

type DailyReportStats struct {
	Day          int64 `json:"day"`
	TotalReports int64 `json:"total_reports"`
}

type MonthlyReportStats struct {
	Month        string             `json:"month"`
	DailyStats   []DailyReportStats `json:"daily_stats"`
	TotalReports int64              `json:"total_reports"`
}

type ReportLittering struct {
	ReportLittering []MonthlyReportStats `json:"report_littering"`
}

type ReportRubbish struct {
	ReportRubbish []MonthlyReportStats `json:"report_rubbish"`
}

type DataUserByAddress struct {
	City      string `json:"city"`
	TotalUser int    `json:"total_user"`
}
