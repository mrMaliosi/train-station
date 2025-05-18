package models

type PositionName struct {
	ID   int    `json:"position_id" db:"position_id"`
	Name string `json:"position_name" db:"position_name"`
}
