package entity

//Course data
type Course struct {
	ID                  ID        `json:"id"`
	Name                string    `json:"course_name"`
	Type                string    `json:"course_type"`
	AverageRating       float64   `json:"course_average_rating"`
	RatedByCount        int       `json:"course_rated_by_count"`
	InstitutionID       int       `json:"institution_id"`
	InstitutionName     string    `json:"institution_name"`
	InstitutionImageURL string    `json:"institution_image_url"`
	MonthlyValueRange   []float64 `json:"course_monthly_value_range"`
	TimeToGraduateRange []int     `json:"course_time_to_graduate_range"`
	Periods             []string  `json:"course_periods"`
}
