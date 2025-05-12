package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type BrigadeRepository interface {
	GetEmployees(ctx context.Context, filter models.BrigadeEmployeeFilter) ([]models.BrigadeEmployee, error)
	CountEmployees(ctx context.Context, filter models.BrigadeEmployeeFilter) (int, error)
}

type brigadeRepository struct {
	db *sqlx.DB
}

func NewBrigadeRepository(db *sqlx.DB) BrigadeRepository {
	return &brigadeRepository{db: db}
}

func (r *brigadeRepository) GetEmployees(ctx context.Context, filter models.BrigadeEmployeeFilter) ([]models.BrigadeEmployee, error) {
	queryBuilder := strings.Builder{}
	args := []interface{}{}

	queryBuilder.WriteString(`
		SELECT 
			e.id, 
			e.name, 
			e.surname, 
			e.patronymic, 
			e.birth_date, 
			e.child_number,
			e.hired_at, 
			e.sex, 
			e.position_id, 
			e.salary,
			b.brigade_id,
			p.position_name AS position_name, 
			EXTRACT(YEAR FROM age(current_date, e.hired_at)) AS experience
		FROM "Brigades" AS b
		JOIN "BrigadeMembers" AS bm ON b.brigade_id = bm.brigade_id
		JOIN "Employees" AS e ON bm.employee_id = e.id
		JOIN "Positions" AS p ON e.position_id = p.position_id
		WHERE 1=1
	`)

	argCounter := 1

	if filter.BrigadeID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND b.brigade_id = $%d", argCounter))
		args = append(args, *filter.BrigadeID)
		argCounter++
	}

	if filter.DepartmentID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND e.position_id IN (SELECT id FROM positions WHERE department_id = $%d)", argCounter))
		args = append(args, *filter.DepartmentID)
		argCounter++
	}

	if filter.LocomotiveID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND b.locomotive_brigade_id = $%d", argCounter))
		args = append(args, *filter.LocomotiveID)
		argCounter++
	}

	if filter.AgeFrom != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.birth_date)) >= $%d", argCounter))
		args = append(args, *filter.AgeFrom)
		argCounter++
	}

	if filter.AgeTo != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.birth_date)) < $%d", argCounter))
		args = append(args, *filter.AgeTo)
		argCounter++
	}

	if filter.SalaryFrom != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND e.salary >= $%d", argCounter))
		args = append(args, *filter.SalaryFrom)
		argCounter++
	}

	if filter.SalaryTo != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND e.salary < $%d", argCounter))
		args = append(args, *filter.SalaryTo)
		argCounter++
	}

	query := queryBuilder.String()

	var employees []models.BrigadeEmployee
	err := r.db.SelectContext(ctx, &employees, query, args...)
	return employees, err
}

func (r *brigadeRepository) CountEmployees(ctx context.Context, filter models.BrigadeEmployeeFilter) (int, error) {
	queryBuilder := strings.Builder{}
	args := []interface{}{}

	queryBuilder.WriteString(`
		SELECT COUNT(*) 
		FROM "Brigades" AS b
		JOIN "BrigadeMembers" AS bm ON b.brigade_id = bm.brigade_id
		JOIN "Employees" AS e ON bm.employee_id = e.id
		WHERE 1=1
	`)

	argCounter := 1

	if filter.BrigadeID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND b.brigade_id = $%d", argCounter))
		args = append(args, *filter.BrigadeID)
		argCounter++
	}

	if filter.DepartmentID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND e.position_id IN (SELECT id FROM positions WHERE department_id = $%d)", argCounter))
		args = append(args, *filter.DepartmentID)
		argCounter++
	}

	if filter.LocomotiveID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND b.locomotive_brigade_id = $%d", argCounter))
		args = append(args, *filter.LocomotiveID)
		argCounter++
	}

	if filter.AgeFrom != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.birth_date)) >= $%d", argCounter))
		args = append(args, *filter.AgeFrom)
		argCounter++
	}

	if filter.AgeTo != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND EXTRACT(YEAR FROM age(current_date, e.birth_date)) < $%d", argCounter))
		args = append(args, *filter.AgeTo)
		argCounter++
	}

	if filter.SalaryFrom != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND e.salary >= $%d", argCounter))
		args = append(args, *filter.SalaryFrom)
		argCounter++
	}

	if filter.SalaryTo != nil {
		queryBuilder.WriteString(fmt.Sprintf(" AND e.salary < $%d", argCounter))
		args = append(args, *filter.SalaryTo)
		argCounter++
	}

	query := queryBuilder.String()

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}
