package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

// Интерфейс для репозитория сотрудников
type EmployeeRepository interface {
	EmployeeFilter(ctx context.Context, f models.EmployeeFilter) ([]models.Employee, error)
	EmployeeCreate(ctx context.Context, f models.EmployeeCreate) (int, error)
	EmployeeDelete(ctx context.Context, id int) error
}

// Реализация репозитория сотрудников
type employeeRepository struct {
	db *sqlx.DB
}

// Конструктор для репозитория сотрудников
func NewEmployeeRepository(db *sqlx.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) EmployeeFilter(
	ctx context.Context,
	f models.EmployeeFilter,
) ([]models.Employee, error) {
	baseQuery := `
    SELECT
        e.id, e.name, e.surname, e.patronymic,
        e.birth_date, e.child_number, e.hired_at,
        e.sex, e.salary,
        e.position_id, d.department_name,
        p.position_name
    FROM "Employees" e
    JOIN "Positions" p ON e.position_id = p.position_id
    JOIN "Departments" d ON d.department_id = p.department_id
    WHERE 1=1`

	var conditions []string
	var args []interface{}
	i := 1

	if f.DepartmentID != nil {
		conditions = append(conditions, fmt.Sprintf("d.department_id = $%d", i))
		args = append(args, *f.DepartmentID)
		i++
	}
	if f.Sex != nil {
		conditions = append(conditions, fmt.Sprintf("e.sex = $%d", i))
		args = append(args, *f.Sex)
		i++
	}
	if f.AgeFrom != nil {
		conditions = append(conditions,
			fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.birth_date)) >= $%d", i))
		args = append(args, *f.AgeFrom)
		i++
	}
	if f.AgeTo != nil {
		conditions = append(conditions,
			fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.birth_date)) < $%d", i))
		args = append(args, *f.AgeTo)
		i++
	}
	if f.ExperienceFrom != nil {
		conditions = append(conditions,
			fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.hired_at)) >= $%d", i))
		args = append(args, *f.ExperienceFrom)
		i++
	}
	if f.ExperienceTo != nil {
		conditions = append(conditions,
			fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.hired_at)) < $%d", i))
		args = append(args, *f.ExperienceTo)
		i++
	}
	if f.ChildrenFrom != nil {
		conditions = append(conditions, fmt.Sprintf("e.child_number >= $%d", i))
		args = append(args, *f.ChildrenFrom)
		i++
	}
	if f.ChildrenTo != nil {
		conditions = append(conditions, fmt.Sprintf("e.child_number < $%d", i))
		args = append(args, *f.ChildrenTo)
		i++
	}
	if f.SalaryFrom != nil {
		conditions = append(conditions, fmt.Sprintf("e.salary >= $%d", i))
		args = append(args, *f.SalaryFrom)
		i++
	}
	if f.SalaryTo != nil {
		conditions = append(conditions, fmt.Sprintf("e.salary < $%d", i))
		args = append(args, *f.SalaryTo)
		i++
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}
	baseQuery += " ORDER BY e.id"

	fmt.Println("Executing query:", baseQuery)
	fmt.Println("With args:", args)

	var employees []models.Employee
	err := r.db.SelectContext(ctx, &employees, baseQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	return employees, nil
}

func (r *employeeRepository) EmployeeCreate(ctx context.Context, req models.EmployeeCreate) (int, error) {
	query := `
    INSERT INTO "Employees" ("id", "name", "surname", "patronymic", "birth_date", "child_number", "hired_at", "sex", "position_id", "salary")
    VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9) 
	RETURNING id`

	var employeeID int
	err := r.db.QueryRow(
		query,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.BirthDate.Time,
		req.ChildNumber,
		req.HiredAt.Time,
		req.Sex,
		req.PositionID,
		req.Salary,
	).Scan(&employeeID)

	if err != nil {
		return -1, fmt.Errorf("query: %w", err)
	}

	return employeeID, nil
}

func (r *employeeRepository) EmployeeDelete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM employees WHERE id = $1", id)
	return err
}
