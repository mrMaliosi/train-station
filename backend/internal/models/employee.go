package models

import (
	"database/sql"
)

type Employee struct {
	ID          int          `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Surname     string       `json:"surname" db:"surname"`
	Patronymic  string       `json:"patronymic" db:"patronymic"`
	BirthDate   sql.NullTime `json:"birth_date" db:"birth_date"` // Используем sql.NullTime для обработки NULL
	ChildNumber int          `json:"child_number" db:"child_number"`
	HiredAt     sql.NullTime `json:"hired_at" db:"hired_at"` // Аналогично для hired_at
	Sex         string       `json:"sex" db:"sex"`
	PositionID  int          `json:"position_id" db:"position_id"`
	Salary      float64      `json:"salary" db:"salary"`
}

type EmployeeFilter struct {
	DepartmentID                 *int
	Sex                          *string
	AgeFrom, AgeTo               *int
	ExperienceFrom, ExperienceTo *int
	ChildrenFrom, ChildrenTo     *int
	SalaryFrom, SalaryTo         *float64
}
