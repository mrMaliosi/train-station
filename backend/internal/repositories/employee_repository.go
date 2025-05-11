package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Employee struct {
	ID         int
	Name       string
	Surname    string
	Patronymic sql.NullString
	BirthDate  time.Time
	ChildNum   int
	HiredAt    time.Time
	Sex        string
	Salary     float64
}

type EmployeeFilter struct {
	DepartmentID                 *int
	Sex                          *string
	AgeFrom, AgeTo               *int
	ExperienceFrom, ExperienceTo *int
	ChildrenFrom, ChildrenTo     *int
	SalaryFrom, SalaryTo         *float64
}

type EmployeeRepository interface {
	Filter(ctx context.Context, f EmployeeFilter) ([]Employee, error)
}

type employeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepo{db: db}
}

// Метод фильтрации работников по различным критериям
func (r *employeeRepo) Filter(ctx context.Context, f EmployeeFilter) ([]Employee, error) {
	query := `
		SELECT e.id, e.name, e.surname, e.patronymic, e.birth_date, e.child_number,
		       e.hired_at, e.sex, e.salary
		FROM "Employees" e
		JOIN "Positions" p ON e.position_id = p.position_id
		JOIN "Departments" d ON p.department_id = d.department_id
		WHERE 1=1`
	args := []interface{}{}
	argID := 1

	// Динамическая подстановка параметров
	if f.DepartmentID != nil {
		query += fmt.Sprintf(" AND d.department_id = $%d", argID)
		args = append(args, *f.DepartmentID)
		argID++
	}
	if f.Sex != nil {
		query += fmt.Sprintf(" AND e.sex = $%d", argID)
		args = append(args, *f.Sex)
		argID++
	}
	if f.AgeFrom != nil {
		query += fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.birth_date)) >= $%d", argID)
		args = append(args, *f.AgeFrom)
		argID++
	}
	if f.AgeTo != nil {
		query += fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.birth_date)) < $%d", argID)
		args = append(args, *f.AgeTo)
		argID++
	}
	if f.ExperienceFrom != nil {
		query += fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.hired_at)) >= $%d", argID)
		args = append(args, *f.ExperienceFrom)
		argID++
	}
	if f.ExperienceTo != nil {
		query += fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.hired_at)) < $%d", argID)
		args = append(args, *f.ExperienceTo)
		argID++
	}
	if f.ChildrenFrom != nil {
		query += fmt.Sprintf(" AND e.child_number >= $%d", argID)
		args = append(args, *f.ChildrenFrom)
		argID++
	}
	if f.ChildrenTo != nil {
		query += fmt.Sprintf(" AND e.child_number < $%d", argID)
		args = append(args, *f.ChildrenTo)
		argID++
	}
	if f.SalaryFrom != nil {
		query += fmt.Sprintf(" AND e.salary >= $%d", argID)
		args = append(args, *f.SalaryFrom)
		argID++
	}
	if f.SalaryTo != nil {
		query += fmt.Sprintf(" AND e.salary < $%d", argID)
		args = append(args, *f.SalaryTo)
		argID++
	}

	query += " ORDER BY e.id"

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		if err := rows.Scan(
			&e.ID, &e.Name, &e.Surname, &e.Patronymic,
			&e.BirthDate, &e.ChildNum, &e.HiredAt, &e.Sex, &e.Salary,
		); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return employees, nil
}
