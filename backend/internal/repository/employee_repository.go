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
}

// Реализация репозитория сотрудников
type employeeRepository struct {
	db *sqlx.DB
}

// Конструктор для репозитория сотрудников
func NewEmployeeRepository(db *sqlx.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) EmployeeFilter(ctx context.Context, f models.EmployeeFilter) ([]models.Employee, error) {
	baseQuery := `
		SELECT 
			e.id, e.name, e.surname, e.patronymic, e.birth_date,
			e.child_number, e.hired_at, e.sex, e.salary
		FROM "Employees" e
		JOIN "Positions" p ON e.position_id = p.position_id
		JOIN "Departments" d ON d.department_id = p.department_id
		WHERE 1=1`

	// Сбор условий и аргументов
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
		conditions = append(conditions, fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.birth_date)) >= $%d", i))
		args = append(args, *f.AgeFrom)
		i++
	}
	if f.AgeTo != nil {
		conditions = append(conditions, fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.birth_date)) < $%d", i))
		args = append(args, *f.AgeTo)
		i++
	}
	if f.ExperienceFrom != nil {
		conditions = append(conditions, fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.hired_at)) >= $%d", i))
		args = append(args, *f.ExperienceFrom)
		i++
	}
	if f.ExperienceTo != nil {
		conditions = append(conditions, fmt.Sprintf("EXTRACT(YEAR FROM age(current_date, e.hired_at)) < $%d", i))
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

	// Собираем полный запрос
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
