package models

import (
	"database/sql"

	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type Employee struct {
	ID             int          `json:"id" db:"id"`
	Name           string       `json:"name" db:"name"`
	Surname        string       `json:"surname" db:"surname"`
	Patronymic     string       `json:"patronymic" db:"patronymic"`
	BirthDate      sql.NullTime `json:"birth_date" db:"birth_date"`
	ChildNumber    int          `json:"child_number" db:"child_number"`
	HiredAt        sql.NullTime `json:"hired_at" db:"hired_at"`
	Sex            string       `json:"sex" db:"sex"`
	PositionID     int          `json:"position_id" db:"position_id"`
	PositionName   string       `json:"position_name" db:"position_name"`
	DepartmentName string       `json:"department_name" db:"department_name"`
	Salary         float64      `json:"salary" db:"salary"`
}

type EmployeeFilter struct {
	DepartmentID                 *int
	Sex                          *string
	AgeFrom, AgeTo               *int
	ExperienceFrom, ExperienceTo *int
	ChildrenFrom, ChildrenTo     *int
	SalaryFrom, SalaryTo         *float64
}

type EmployeeCreate struct {
	Name        string             `json:"name" db:"name"`
	Surname     string             `json:"surname" db:"surname"`
	Patronymic  *string            `json:"patronymic,omitempty" db:"patronymic"`
	BirthDate   utilities.DateOnly `json:"birth_date" db:"birth_date"`
	ChildNumber int                `json:"child_number" db:"child_number"`
	HiredAt     utilities.DateOnly `json:"hired_at" db:"hired_at"`
	Sex         string             `json:"sex" db:"sex"`
	PositionID  int                `json:"position_id" db:"position_id"`
	Salary      float64            `json:"salary" db:"salary"`
}
