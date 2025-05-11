package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

type TicketRepository interface {
	GetSoldTickets(ctx context.Context, filter models.TicketFilter) ([]models.Ticket, error)
	CountSoldTickets(ctx context.Context, filter models.TicketFilter) (int, error)
	GetUnsoldTickets(ctx context.Context, filter models.TicketFilter) ([]models.Ticket, error)
	CountReturnedTickets(ctx context.Context, filter models.TicketFilter) (int, error)
}

type ticketRepository struct {
	db *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) GetSoldTickets(ctx context.Context, filter models.TicketFilter) ([]models.Ticket, error) {
	query := `
	SELECT ti.*
	FROM "Tickets" ti
	JOIN "Routes" r ON ti.route_id = r.route_id
	WHERE ti.ticket_status = 'куплен'`
	args := []interface{}{}

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

	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var tickets []models.Ticket
	err := r.db.SelectContext(ctx, &tickets, query, args...)
	return tickets, err
}

func (r *ticketRepository) CountSoldTickets(ctx context.Context, filter models.TicketFilter) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM "Tickets" ti
	JOIN "Routes" r ON ti.route_id = r.route_id
	WHERE ti.ticket_status = 'куплен'`
	args := []interface{}{}

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

	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}

func (r *ticketRepository) GetUnsoldTickets(ctx context.Context, filter models.TicketFilter) ([]models.Ticket, error) {
	query := `
	SELECT ti.*
	FROM "Tickets" ti
	JOIN "Routes" r ON ti.route_id = r.route_id
	WHERE ti.ticket_status = 'свободен'`
	args := []interface{}{}

	if filter.RouteID != nil {
		query += " AND r.route_id = ?"
		args = append(args, *filter.RouteID)
	}
	if filter.Date != nil {
		query += " AND DATE(r.start_time) = ?"
		args = append(args, *filter.Date)
	}

	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var tickets []models.Ticket
	err := r.db.SelectContext(ctx, &tickets, query, args...)
	return tickets, err
}

func (r *ticketRepository) CountReturnedTickets(ctx context.Context, filter models.TicketFilter) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM "Tickets" ti
	JOIN "Routes" r ON ti.route_id = r.route_id
	WHERE ti.ticket_status = 'возвращён'`
	args := []interface{}{}

	if filter.RouteID != nil {
		query += " AND r.route_id = ?"
		args = append(args, *filter.RouteID)
	}
	if filter.Date != nil {
		query += " AND DATE(r.start_time) = ?"
		args = append(args, *filter.Date)
	}

	query = strings.ReplaceAll(query, "?", "$%d")
	for i := range args {
		query = strings.Replace(query, "$%d", fmt.Sprintf("$%d", i+1), 1)
	}

	var count int
	err := r.db.GetContext(ctx, &count, query, args...)
	return count, err
}
