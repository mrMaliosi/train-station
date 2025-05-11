package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type RouteRepository interface {
	GetFilteredRoutes(ctx context.Context, filter models.RouteFilter) ([]models.Route, error)
	CountFilteredRoutes(ctx context.Context, filter models.RouteFilter) (int, error)
	GetReturnedTicketsDuringDelay(ctx context.Context, filter models.RouteFilter) (int, error)
}

type routeRepository struct {
	db *sqlx.DB
}

func NewRouteRepository(db *sqlx.DB) RouteRepository {
	return &routeRepository{db: db}
}

func (r *routeRepository) GetFilteredRoutes(ctx context.Context, filter models.RouteFilter) ([]models.Route, error) {
	query := `
	SELECT DISTINCT r.*
	FROM "Routes" AS r
	LEFT JOIN "DelayReason" AS dr ON r.route_id = dr.route_id
	LEFT JOIN "Trains" AS t ON r.train_number = t.train_number
	LEFT JOIN "RoutesStations" AS rs ON r.route_id = rs.route_id
	LEFT JOIN "Stations" AS s ON rs.station_id = s.station_id
	WHERE 1=1
	`

	args := []interface{}{}
	if filter.RouteID != nil {
		query += " AND r.route_id = ?"
		args = append(args, *filter.RouteID)
	}
	if filter.Status != nil {
		query += " AND r.status = ?"
		args = append(args, *filter.Status)
	}
	if filter.DelayReason != nil {
		query += " AND dr.reason = ?"
		args = append(args, *filter.DelayReason)
	}
	if filter.TrainType != nil {
		query += " AND t.train_type = ?"
		args = append(args, *filter.TrainType)
	}
	if filter.StationName != nil {
		query += " AND s.station_name = ? AND rs.station_number > 1"
		args = append(args, *filter.StationName)
	}

	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var routes []models.Route
	err := r.db.SelectContext(ctx, &routes, query, args...)
	return routes, err
}

func (r *routeRepository) CountFilteredRoutes(ctx context.Context, filter models.RouteFilter) (int, error) {
	query := `
	SELECT COUNT(DISTINCT r.route_id)
	FROM "Routes" AS r
	LEFT JOIN "DelayReason" AS dr ON r.route_id = dr.route_id
	LEFT JOIN "Trains" AS t ON r.train_number = t.train_number
	LEFT JOIN "RoutesStations" AS rs ON r.route_id = rs.route_id
	LEFT JOIN "Stations" AS s ON rs.station_id = s.station_id
	WHERE 1=1
	`

	args := []interface{}{}
	if filter.RouteID != nil {
		query += " AND r.route_id = ?"
		args = append(args, *filter.RouteID)
	}
	if filter.Status != nil {
		query += " AND r.status = ?"
		args = append(args, *filter.Status)
	}
	if filter.DelayReason != nil {
		query += " AND dr.reason = ?"
		args = append(args, *filter.DelayReason)
	}
	if filter.TrainType != nil {
		query += " AND t.train_type = ?"
		args = append(args, *filter.TrainType)
	}
	if filter.StationName != nil {
		query += " AND s.station_name = ? AND rs.station_number > 1"
		args = append(args, *filter.StationName)
	}

	// PostgreSQL-style подстановки
	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}

func (r *routeRepository) GetReturnedTicketsDuringDelay(ctx context.Context, filter models.RouteFilter) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM "Routes" AS r
	JOIN "Tickets" AS t ON r.route_id = t.route_id
	JOIN "DelayReason" AS dr ON r.route_id = dr.route_id
	WHERE t.ticket_status = 'возвращён'
	`

	args := []interface{}{}
	if filter.RouteID != nil {
		query += " AND r.route_id = ?"
		args = append(args, *filter.RouteID)
	}
	if filter.DelayReason != nil {
		query += " AND dr.reason = ?"
		args = append(args, *filter.DelayReason)
	}

	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}
