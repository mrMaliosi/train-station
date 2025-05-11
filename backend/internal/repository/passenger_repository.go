package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type PassengerRepository struct {
	db *sqlx.DB
}

func NewPassengerRepository(db *sqlx.DB) *PassengerRepository {
	return &PassengerRepository{db: db}
}

func (r *PassengerRepository) GetFilteredPassengers(filter models.PassengerFilter) ([]models.PassengerWithInfo, error) {
	query := `SELECT DISTINCT ON (p.passenger_id)
		p.passenger_id,
		p.name,
		p.surname,
		p.patronimic,
		p.sex,
		p.birth_date,
		EXTRACT(YEAR FROM age(current_date, p.birth_date)) AS age,
		rts.route_id,
		DATE(art.real_arrival_time) AS travel_date,
		lu.luggage_id IS NOT NULL AS has_luggage
	FROM "Passengers" AS p
	JOIN "Tickets" AS ti ON ti.passenger_id = p.passenger_id
	LEFT JOIN "TicketsLuggage" AS tl ON ti.ticket_id = tl.ticket_id
	LEFT JOIN "Luggage" AS lu ON tl.luggage_id = lu.luggage_id
	JOIN "Routes" AS rts ON rts.route_id = ti.route_id
	JOIN "RoutesStations" AS rs ON rs.route_id = rts.route_id
	JOIN "ArrivalTime" AS art ON rs.arrival_time_id = art.arrival_id
	JOIN "Stations" AS s ON s.station_id = rs.station_id`

	var conditions []string
	params := map[string]interface{}{}

	// Применение фильтров
	if filter.RouteID != nil {
		conditions = append(conditions, "rts.route_id = :route_id")
		params["route_id"] = *filter.RouteID
	}
	if filter.TravelDate != nil {
		conditions = append(conditions, "DATE(art.real_arrival_time) = :travel_date")
		params["travel_date"] = *filter.TravelDate
	}
	if filter.Abroad != nil {
		conditions = append(conditions, "s.is_abroad = :is_abroad")
		params["is_abroad"] = *filter.Abroad
	}
	if filter.HasLuggage != nil {
		if *filter.HasLuggage {
			conditions = append(conditions, "lu.luggage_id IS NOT NULL")
		} else {
			conditions = append(conditions, "lu.luggage_id IS NULL")
		}
	}
	if filter.Sex != nil {
		conditions = append(conditions, "p.sex = :sex")
		params["sex"] = *filter.Sex
	}
	if filter.MinAge != nil {
		conditions = append(conditions, "EXTRACT(YEAR FROM age(current_date, p.birth_date)) >= :min_age")
		params["min_age"] = *filter.MinAge
	}
	if filter.MaxAge != nil {
		conditions = append(conditions, "EXTRACT(YEAR FROM age(current_date, p.birth_date)) < :max_age")
		params["max_age"] = *filter.MaxAge
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	namedStmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, fmt.Errorf("prepare query: %w", err)
	}
	defer namedStmt.Close()

	var passengers []models.PassengerWithInfo
	err = namedStmt.Select(&passengers, params)
	if err != nil {
		return nil, fmt.Errorf("select passengers: %w", err)
	}

	return passengers, nil
}
