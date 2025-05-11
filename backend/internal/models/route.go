package models

type Route struct {
	RouteID           int     `db:"route_id" json:"route_id"`
	TrainNumber       string  `db:"train_number" json:"train_number"`
	StartTime         string  `db:"start_time" json:"start_time"`
	EndTime           string  `db:"end_time" json:"end_time"`
	RealArrivalTime   *string `db:"real_arrival_time" json:"real_arrival_time,omitempty"`
	RealDepartureTime *string `db:"real_departure_time" json:"real_departure_time,omitempty"`
	Status            *string `db:"status" json:"status,omitempty"`
}

type RouteFilter struct {
	RouteID     *int    // ID маршрута
	Status      *string // Статус маршрута (например, "отменён", "задерживается")
	DelayReason *string // Причина задержки
	TrainType   *string // Тип поезда ("пригородный" и т.п.)
	StationName *string // Название станции (например, "Грасиона")
}
