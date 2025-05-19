package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

// Интерфейс для репозитория сотрудников
type DepartmentRepository interface {
	DepartmentsSelect(ctx context.Context) ([]models.DepartmentName, error)
	DepartmentsInfo(ctx context.Context) ([]models.DepartmentInfo, error)
}

// Реализация репозитория сотрудников
type departmentRepository struct {
	db *sqlx.DB
}

// Конструктор для репозитория сотрудников
func NewDepartmentRepository(db *sqlx.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}

func (r *departmentRepository) DepartmentsSelect(
	ctx context.Context,
) ([]models.DepartmentName, error) {
	baseQuery := `
    SELECT d.department_id, d.department_name
	FROM "Departments" AS d`

	fmt.Println("Executing query:", baseQuery)

	var departments []models.DepartmentName
	err := r.db.SelectContext(ctx, &departments, baseQuery)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	return departments, nil
}

func (r *departmentRepository) DepartmentsInfo(
	ctx context.Context,
) ([]models.DepartmentInfo, error) {
	baseQuery := `
		SELECT 
			d.department_name,
			COALESCE(e.name, '') AS name,
			COALESCE(e.surname, '') AS surname,
			COALESCE(e.patronymic, '') AS patronymic,
			e.birth_date
		FROM "Departments" AS d
		LEFT JOIN "Employees" AS e ON d.director_id = e.id
	`

	fmt.Println("Executing query:", baseQuery)

	var departments []models.DepartmentInfo
	err := r.db.SelectContext(ctx, &departments, baseQuery)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	return departments, nil
}
