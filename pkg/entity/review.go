package entity

import "time"

//Review entity
type Review struct {
	ID            ID      `json:"id"`
	StudentID     uint    `json:"student_id"`
	InstitutionID uint    `json:"institution_id"`
	Rate          float64 `json:"rate"`
	Title         string  `json:"title"`
	Pros          string  `json:"pros"`
	Cons          string  `json:"cons"`
	Suggestion    string  `json:"suggestion"`
	CourseID      uint    `json:"course_id"`
	CourseInfo    struct {
		CourseID            uint   `json:"course_id"`
		CourseType          string `json:"course_type"`
		Period              string `json:"period"`
		StartYear           int    `json:"start_year"`
		CourseName          string `json:"course_name"`
		MonthlyPaymentValue int    `json:"monthly_payment_value"`
	} `json:"course_info"`
	CreatedAt time.Time `json:"created_at"`
}
