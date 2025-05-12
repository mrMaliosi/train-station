package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type TicketRepository interface {
	GetTickets(ctx context.Context, filter models.TicketFilter) ([]models.Ticket, error)
	CountSoldTickets(ctx context.Context, filter models.TicketFilter) (int, error)
	CountReturnedTickets(ctx context.Context, filter models.TicketFilter) (int, error)
	GetTicketStatuses(ctx context.Context) ([]string, error)
}

type ticketRepository struct {
	db *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) TicketRepository {
	return &ticketRepository{db: db}
}

// Обобщённый метод получения билетов по фильтру
func (r *ticketRepository) GetTickets(ctx context.Context, filter models.TicketFilter) ([]models.Ticket, error) {
	query := `
    SELECT 
        ti.ticket_id,
        ti.route_id,
        ti.ticket_status,
        ti.passenger_id,
        ti.bought_at,
        ti.price,
        r.train_number  -- Номер поезда
    FROM 
        "Tickets" ti
    JOIN 
        "Routes" r ON ti.route_id = r.route_id
    WHERE 1 = 1`
	args := []interface{}{}

	if filter.Status != nil {
		query += " AND ti.ticket_status = ?"
		args = append(args, *filter.Status)
	}
	if filter.FromDate != nil {
		query += " AND ti.bought_at >= ?"
		args = append(args, *filter.FromDate)
	}
	if filter.ToDate != nil {
		query += " AND ti.bought_at < ?"
		args = append(args, *filter.ToDate)
	}
	if filter.RouteID != nil {
		query += " AND r.route_id = ?"
		args = append(args, *filter.RouteID)
	}

	query = preparePostgresQuery(query, args)

	var tickets []models.Ticket
	err := r.db.SelectContext(ctx, &tickets, query, args...)
	return tickets, err
}

// Подсчёт купленных билетов
func (r *ticketRepository) CountSoldTickets(ctx context.Context, filter models.TicketFilter) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM "Tickets" ti
		JOIN "Routes" r ON ti.route_id = r.route_id
		WHERE ti.ticket_status = 'куплен'` // Статус для проданных билетов
	args := []interface{}{}

	// Добавьте фильтры, если они присутствуют
	if filter.RouteID != nil {
		query += " AND r.route_id = $1"
		args = append(args, *filter.RouteID)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}

// Подсчёт возвращённых билетов
// Пример для подсчёта возвращённых билетов
func (r *ticketRepository) CountReturnedTickets(ctx context.Context, filter models.TicketFilter) (int, error) {
	query := `
		SELECT COUNT(*)
		FROM "Tickets" ti
		JOIN "Routes" r ON ti.route_id = r.route_id
		WHERE ti.ticket_status = 'возвращён'` // Статус для возвращённых билетов
	args := []interface{}{}

	// Добавьте фильтры, если они присутствуют
	if filter.RouteID != nil {
		query += " AND r.route_id = $1"
		args = append(args, *filter.RouteID)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}

// Получение возможных значений ENUM ticket_status
func (r *ticketRepository) GetTicketStatuses(ctx context.Context) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT unnest(enum_range(NULL::"TICKET_STATUS"))`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []string
	for rows.Next() {
		var status string
		if err := rows.Scan(&status); err == nil {
			statuses = append(statuses, status)
		}
	}
	return statuses, nil
}

// Вспомогательная функция замены ? на $1, $2 и т.д.
func preparePostgresQuery(query string, args []interface{}) string {
	count := 0
	return strings.ReplaceAll(query, "?", func() string {
		count++
		return fmt.Sprintf("$%d", count)
	}())
}
