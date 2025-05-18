package models

type DepartmentName struct {
	ID   int    `json:"id" db:"department_id"`
	Name string `json:"department_name" db:"department_name"`
}
