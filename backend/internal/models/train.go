package models

import "time"

// Train представляет данные о поезде
type Train struct {
	TrainNumber int       `json:"train_number"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	RouteTime   int       `json:"route_time"` // разница между временем начала и конца маршрута
	Price       float64   `json:"price"`      // цена билета
}

// TrainFilter представляет фильтры для поиска поездов
type TrainFilter struct {
	RouteID   *int     // ID маршрута
	PriceMin  *float64 // Минимальная цена
	PriceMax  *float64 // Максимальная цена
	RouteTime *int     // Длительность маршрута
}
