package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type LocomotiveDriverRepository interface {
	GetLocomotiveDrivers(ctx context.Context, f models.LocomotiveDriverFilter) ([]models.Employee, error)
	GetLocomotiveDriversCount(ctx context.Context, f models.LocomotiveDriverFilter) (int, error)
}

type locomotiveDriverRepository struct {
	db *sqlx.DB
}

func NewLocomotiveDriverRepository(db *sqlx.DB) LocomotiveDriverRepository {
	return &locomotiveDriverRepository{db: db}
}

func (r *locomotiveDriverRepository) GetLocomotiveDrivers(ctx context.Context, f models.LocomotiveDriverFilter) ([]models.Employee, error) {
	baseQuery := `
		SELECT DISTINCT ON (e.id) e.*
		FROM "Employees" AS e
		JOIN "Positions" AS p ON p.position_id = e.position_id
		JOIN "MedicalCheckups" AS m ON m.employee_id = e.id
		WHERE p.position_name = 'Водитель локомотива'`

	var conditions []string
	var args []interface{}
	i := 1

	// Фильтрация по году медосмотра
	if f.MedicalCheckupYear != nil {
		conditions = append(conditions, fmt.Sprintf("EXTRACT(YEAR FROM m.medical_checkup_date) = $%d", i))
		args = append(args, *f.MedicalCheckupYear)
		i++
	}

	// Фильтрация по полу
	if f.Sex != nil {
		conditions = append(conditions, fmt.Sprintf("e.sex = $%d", i))
		args = append(args, *f.Sex)
		i++
	}

	// Фильтрация по возрасту
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

	// Фильтрация по зарплате
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

	var drivers []models.Employee
	err := r.db.SelectContext(ctx, &drivers, baseQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return drivers, nil
}

func (r *locomotiveDriverRepository) GetLocomotiveDriversCount(ctx context.Context, f models.LocomotiveDriverFilter) (int, error) {
	baseQuery := `
		SELECT COUNT(*) 
		FROM (
			SELECT DISTINCT ON (e.id) e.*
			FROM "Employees" AS e
			JOIN "Positions" AS p ON p.position_id = e.position_id
			JOIN "MedicalCheckups" AS m ON m.employee_id = e.id
			WHERE p.position_name = 'Водитель локомотива'`

	var conditions []string
	var args []interface{}
	i := 1

	// Фильтрация по году медосмотра
	if f.MedicalCheckupYear != nil {
		conditions = append(conditions, fmt.Sprintf("EXTRACT(YEAR FROM m.medical_checkup_date) = $%d", i))
		args = append(args, *f.MedicalCheckupYear)
		i++
	}

	// Фильтрация по полу
	if f.Sex != nil {
		conditions = append(conditions, fmt.Sprintf("e.sex = $%d", i))
		args = append(args, *f.Sex)
		i++
	}

	// Фильтрация по возрасту
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

	// Фильтрация по зарплате
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
	baseQuery += " ) AS sub"

	var count int
	err := r.db.GetContext(ctx, &count, baseQuery, args...)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}

	return count, nil
}
