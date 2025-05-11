package models

type Ticket struct {
	TicketID     int     `db:"ticket_id" json:"ticket_id"`
	RouteID      int     `db:"route_id" json:"route_id"`
	TicketStatus string  `db:"ticket_status" json:"ticket_status"`
	PassengerID  int     `db:"passenger_id" json:"passenger_id"`
	BoughtAt     *string `db:"bought_at" json:"bought_at,omitempty"`
	Price        float64 `db:"price" json:"price"`
}

type TicketFilter struct {
	RouteID  *int
	Date     *string
	FromDate *string
	ToDate   *string
}
