package models

import (
	"database/sql"
)

// Ticket представляет информацию о билете.
type Ticket struct {
	TicketID     int          `db:"ticket_id"`
	RouteID      int          `db:"route_id"`
	TicketStatus string       `db:"ticket_status"`
	PassengerID  int          `db:"passenger_id"`
	BoughtAt     sql.NullTime `db:"bought_at"`
	Price        float64      `db:"price"`
	TrainNumber  string       `db:"train_number"` // Добавлено поле для номера поезда
}

type TicketFilter struct {
	RouteID  *int    `form:"routeID"`
	FromDate *string `form:"fromDate"`
	ToDate   *string `form:"toDate"`
	Status   *string `form:"status"`
}
