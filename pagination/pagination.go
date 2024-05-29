package pagination

type Pageinfo struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
}

type CountDataInfo struct {
	TotalCount         int `json:"total_count"`
	CountPerluDitinjau int `json:"count_pending"`
	CountDiterima      int `json:"count_approved"`
	CountDitolak       int `json:"count_rejected"`
}
