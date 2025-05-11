package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type TrainRepository interface {
	GetTrains(ctx context.Context, filter models.TrainFilter) ([]models.Train, error)
	GetTrainsCount(ctx context.Context, filter models.TrainFilter) (int, error)
}

type trainRepository struct {
	db *sqlx.DB
}

func NewTrainRepository(db *sqlx.DB) TrainRepository {
	return &trainRepository{db: db}
}

// GetTrains реализует получение списка поездов с фильтрацией
func (r *trainRepository) GetTrains(ctx context.Context, filter models.TrainFilter) ([]models.Train, error) {
	query := `
		SELECT t.train_number, r.start_time, r.end_time, 
		       EXTRACT(EPOCH FROM (r.end_time - r.start_time)) / 60 AS route_time, 
		       ti.price
		FROM "Trains" AS t
		JOIN "Routes" AS r ON r.train_number = t.train_number
		LEFT JOIN "Tickets" AS ti ON ti.route_id = r.route_id
		WHERE 1=1
	`
	args := []interface{}{}

	if filter.RouteID != nil {
		query += ` AND r.route_id = $1`
		args = append(args, *filter.RouteID)
	}
	if filter.PriceMin != nil {
		query += ` AND ti.price >= $2`
		args = append(args, *filter.PriceMin)
	}
	if filter.PriceMax != nil {
		query += ` AND ti.price <= $3`
		args = append(args, *filter.PriceMax)
	}
	if filter.RouteTime != nil {
		query += ` AND EXTRACT(EPOCH FROM (r.end_time - r.start_time)) / 60 <= $4`
		args = append(args, *filter.RouteTime)
	}

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trains []models.Train
	for rows.Next() {
		var train models.Train
		if err := rows.StructScan(&train); err != nil {
			return nil, err
		}
		trains = append(trains, train)
	}

	return trains, nil
}

// GetTrainsCount возвращает количество поездов, удовлетворяющих фильтрам
func (r *trainRepository) GetTrainsCount(ctx context.Context, filter models.TrainFilter) (int, error) {
	query := `
		SELECT COUNT(*) 
		FROM "Trains" AS t
		JOIN "Routes" AS r ON r.train_number = t.train_number
		LEFT JOIN "Tickets" AS ti ON ti.route_id = r.route_id
		WHERE 1=1
	`
	args := []interface{}{}

	if filter.RouteID != nil {
		query += ` AND r.route_id = $1`
		args = append(args, *filter.RouteID)
	}
	if filter.PriceMin != nil {
		query += ` AND ti.price >= $2`
		args = append(args, *filter.PriceMin)
	}
	if filter.PriceMax != nil {
		query += ` AND ti.price <= $3`
		args = append(args, *filter.PriceMax)
	}
	if filter.RouteTime != nil {
		query += ` AND EXTRACT(EPOCH FROM (r.end_time - r.start_time)) / 60 <= $4`
		args = append(args, *filter.RouteTime)
	}

	var count int
	if err := r.db.QueryRowxContext(ctx, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
