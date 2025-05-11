package models

import (
	"database/sql"
)

type BrigadeEmployee struct {
	ID          int          `db:"id"`
	Name        string       `db:"name"`
	Surname     string       `db:"surname"`
	Patronymic  string       `db:"patronymic"`
	BirthDate   sql.NullTime `db:"birth_date"`
	ChildNumber int          `db:"child_number"`
	HiredAt     sql.NullTime `db:"hired_at"`
	Sex         string       `db:"sex"`
	PositionID  int          `db:"position_id"`
	Salary      float64      `db:"salary"`

	PositionName string `db:"position_name"`
	Experience   int    `db:"experience"`
}

type BrigadeEmployeeFilter struct {
	BrigadeID    *int `form:"brigade_id"`
	DepartmentID *int `form:"department_id"`
	LocomotiveID *int `form:"locomotive_id"`
	AgeFrom      *int `form:"age_from"`
	AgeTo        *int `form:"age_to"`
	SalaryFrom   *int `form:"salary_from"`
	SalaryTo     *int `form:"salary_to"`
}

type BrigadeEmployeesResponse struct {
	Employees []Employee `json:"employees"`
	Count     int        `json:"count"`
}
