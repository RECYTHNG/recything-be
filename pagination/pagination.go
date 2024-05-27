package pagination

type Pageinfo struct {
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	LastPage int `json:"lastPage"`
}
