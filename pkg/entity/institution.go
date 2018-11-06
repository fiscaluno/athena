package entity

//Institution data
type Institution struct {
	ID            ID       `json:"id"`
	Name          string   `json:"name"`
	ImageURL      string   `json:"image_url"`
	AverageRating float64  `json:"average_rating"`
	RatedByCount  int      `json:"rated_by_count"`
	Website       string   `json:"website"`
	Cnpj          string   `json:"cnpj"`
	Address       string   `json:"address"`
	City          string   `json:"city"`
	Province      string   `json:"province"`
	Emails        []string `json:"emails"`
	Phones        []string `json:"phones"`
}
