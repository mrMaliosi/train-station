package models

import "database/sql"

type DepartmentName struct {
	ID   int    `json:"id" db:"department_id"`
	Name string `json:"department_name" db:"department_name"`
}

type DepartmentInfo struct {
	Name               string       `json:"department_name" db:"department_name"`
	DirectorName       string       `json:"director_name" db:"name"`
	DirectorSurname    string       `json:"director_surname" db:"surname"`
	DirectorPatronymic string       `json:"director_patronymic" db:"patronymic"`
	BirthDate          sql.NullTime `json:"birth_date" db:"birth_date"`
}
