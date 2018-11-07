package entity

// DetailedReview data
type DetailedReview struct {
	ID              ID   `json:"id"`
	ReviewID        uint `json:"review_id"`
	InstitutionID   uint `json:"institution_id"`
	CourseID        uint `json:"course_id"`
	DetailedReviews []struct {
		ReviewType int     `json:"review_type"`
		Rate       float64 `json:"rate"`
	} `json:"detailed_reviews"`
}
