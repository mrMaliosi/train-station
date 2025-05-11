package models

type LocomotiveDriverFilter struct {
	Sex                  *string  `json:"sex,omitempty"`
	AgeFrom, AgeTo       *int     `json:"age_from,omitempty,age_to,omitempty"`
	SalaryFrom, SalaryTo *float64 `json:"salary_from,omitempty,salary_to,omitempty"`
	MedicalCheckupYear   *int     `json:"medical_checkup_year,omitempty"`
}
