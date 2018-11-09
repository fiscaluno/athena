package entity

//Institution data
type Institution struct {
	ID            ID       `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ImageURL      string   `json:"image_url,omitempty"`
	AverageRating float64  `json:"average_rating,omitempty"`
	RatedByCount  int      `json:"rated_by_count,omitempty"`
	Website       string   `json:"website,omitempty"`
	Cnpj          string   `json:"cnpj,omitempty"`
	Address       string   `json:"address,omitempty"`
	City          string   `json:"city,omitempty"`
	Province      string   `json:"province,omitempty"`
	Emails        []string `json:"emails,omitempty"`
	Phones        []string `json:"phones,omitempty"`
}
