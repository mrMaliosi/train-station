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
