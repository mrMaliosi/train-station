package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type LocomotiveRepository interface {
	GetLocomotives(ctx context.Context, filter models.LocomotiveFilter) ([]models.Locomotive, error)
	GetLocomotivesCount(ctx context.Context, filter models.LocomotiveFilter) (int, error)
}

type locomotiveRepository struct {
	db *sqlx.DB
}

func NewLocomotiveRepository(db *sqlx.DB) LocomotiveRepository {
	return &locomotiveRepository{db: db}
}

func (r *locomotiveRepository) GetLocomotives(ctx context.Context, filter models.LocomotiveFilter) ([]models.Locomotive, error) {
	query := `
		SELECT DISTINCT ON (l.id) l.*
		FROM "Locomotives" AS l
		JOIN "Repairs" AS r ON l.id = r.locomotive_id
		WHERE 1=1
	`

	// Фильтры по дате ремонта
	if filter.RepairStartDateMin != nil {
		query += " AND r.repair_start_date >= $1"
	}
	if filter.RepairStartDateMax != nil {
		query += " AND r.repair_start_date <= $2"
	}
	if filter.RepairEndDateMin != nil {
		query += " AND r.repair_end_date >= $3"
	}
	if filter.RepairEndDateMax != nil {
		query += " AND r.repair_end_date <= $4"
	}
	if filter.RepairType != nil {
		query += " AND r.repair_type = $5"
	}

	// Фильтры по количеству ремонтов
	if filter.RepairCountMin != nil {
		query += " AND (SELECT COUNT(*) FROM \"Repairs\" r2 WHERE r2.locomotive_id = l.id) >= $6"
	}
	if filter.RepairCountMax != nil {
		query += " AND (SELECT COUNT(*) FROM \"Repairs\" r2 WHERE r2.locomotive_id = l.id) <= $7"
	}

	// Фильтры по возрасту
	if filter.AgeMin != nil {
		query += " AND (EXTRACT(YEAR FROM age(l.put_into_service)) >= $8)"
	}
	if filter.AgeMax != nil {
		query += " AND (EXTRACT(YEAR FROM age(l.put_into_service)) <= $9)"
	}

	// Добавление лимита по станции и другие фильтры
	if filter.StationID != nil {
		query += " AND l.base_station_id = $10"
	}

	// Исполнение запроса
	rows, err := r.db.QueryContext(ctx, query, filter.RepairStartDateMin, filter.RepairStartDateMax, filter.RepairEndDateMin, filter.RepairEndDateMax, filter.RepairType, filter.RepairCountMin, filter.RepairCountMax, filter.AgeMin, filter.AgeMax, filter.StationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locomotives []models.Locomotive
	for rows.Next() {
		var l models.Locomotive
		if err := rows.Scan(&l.ID, &l.Model, &l.Status, &l.LocomotiveBrigadeID, &l.TechnicBrigadeID, &l.PutIntoService, &l.BaseStationID); err != nil {
			return nil, err
		}
		locomotives = append(locomotives, l)
	}

	return locomotives, nil
}

func (r *locomotiveRepository) GetLocomotivesCount(ctx context.Context, filter models.LocomotiveFilter) (int, error) {
	query := `
		SELECT COUNT(DISTINCT l.id)
		FROM "Locomotives" AS l
		JOIN "Repairs" AS r ON l.id = r.locomotive_id
		WHERE 1=1
	`

	// Фильтры по дате ремонта
	if filter.RepairStartDateMin != nil {
		query += " AND r.repair_start_date >= $1"
	}
	if filter.RepairStartDateMax != nil {
		query += " AND r.repair_start_date <= $2"
	}
	if filter.RepairEndDateMin != nil {
		query += " AND r.repair_end_date >= $3"
	}
	if filter.RepairEndDateMax != nil {
		query += " AND r.repair_end_date <= $4"
	}
	if filter.RepairType != nil {
		query += " AND r.repair_type = $5"
	}

	// Фильтры по количеству ремонтов
	if filter.RepairCountMin != nil {
		query += " AND (SELECT COUNT(*) FROM \"Repairs\" r2 WHERE r2.locomotive_id = l.id) >= $6"
	}
	if filter.RepairCountMax != nil {
		query += " AND (SELECT COUNT(*) FROM \"Repairs\" r2 WHERE r2.locomotive_id = l.id) <= $7"
	}

	// Фильтры по возрасту
	if filter.AgeMin != nil {
		query += " AND (EXTRACT(YEAR FROM age(l.put_into_service)) >= $8)"
	}
	if filter.AgeMax != nil {
		query += " AND (EXTRACT(YEAR FROM age(l.put_into_service)) <= $9)"
	}

	// Добавление лимита по станции и другие фильтры
	if filter.StationID != nil {
		query += " AND l.base_station_id = $10"
	}

	// Исполнение запроса
	var count int
	err := r.db.QueryRowContext(ctx, query, filter.RepairStartDateMin, filter.RepairStartDateMax, filter.RepairEndDateMin, filter.RepairEndDateMax, filter.RepairType, filter.RepairCountMin, filter.RepairCountMax, filter.AgeMin, filter.AgeMax, filter.StationID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
