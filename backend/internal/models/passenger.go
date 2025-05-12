package models

import "time"

// Passenger представляет информацию о пассажире.
type Passenger struct {
	PassengerID int       `db:"passenger_id" json:"passenger_id"`
	Name        string    `db:"name" json:"name"`
	Surname     string    `db:"surname" json:"surname"`
	Patronimic  *string   `db:"patronimic" json:"patronimic,omitempty"`
	Sex         string    `db:"sex" json:"sex"` // 'M' или 'F'
	BirthDate   time.Time `db:"birth_date" json:"birth_date"`
}

type PassengerWithInfo struct {
	PassengerID int       `db:"passenger_id" json:"passenger_id"`
	Name        string    `db:"name" json:"name"`
	Surname     string    `db:"surname" json:"surname"`
	Patronimic  *string   `db:"patronimic" json:"patronimic,omitempty"`
	Sex         string    `db:"sex" json:"sex"`
	BirthDate   time.Time `db:"birth_date" json:"birth_date"`

	Age        int    `db:"age" json:"age"`
	RouteID    int    `db:"route_id" json:"route_id"`
	TravelDate string `db:"travel_date" json:"travel_date"`
	HasLuggage bool   `db:"has_luggage" json:"has_luggage"`
}

// PassengerFilter используется для фильтрации пассажиров по различным признакам.
type PassengerFilter struct {
	RouteID    *int    `form:"routeID"`
	Sex        *string `form:"sex"`
	MinAge     *int    `form:"minAge"`
	MaxAge     *int    `form:"maxAge"`
	HasLuggage *bool   `form:"hasLuggage"`
	Abroad     *bool   `form:"abroad"`
	TravelDate *string `form:"travelDate"`
}
