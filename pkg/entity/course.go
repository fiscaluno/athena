package entity

//Course data
type Course struct {
	ID                  ID        `json:"course_id,omitempty"`
	Name                string    `json:"course_name,omitempty"`
	Type                string    `json:"course_type,omitempty"`
	AverageRating       float64   `json:"course_average_rating,omitempty"`
	RatedByCount        int       `json:"course_rated_by_count,omitempty"`
	InstitutionID       int       `json:"institution_id,omitempty"`
	InstitutionName     string    `json:"institution_name,omitempty"`
	InstitutionImageURL string    `json:"institution_image_url,omitempty"`
	MonthlyValueRange   []float64 `json:"course_monthly_value_range,omitempty"`
	TimeToGraduateRange []int     `json:"course_time_to_graduate_range,omitempty"`
	Periods             []string  `json:"course_periods,omitempty"`
}
