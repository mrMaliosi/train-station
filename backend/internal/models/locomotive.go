package models

import "time"

type Locomotive struct {
	ID                  int        `json:"id"`
	Model               string     `json:"model"`
	Status              string     `json:"status"`
	LocomotiveBrigadeID int        `json:"locomotive_brigade_id"`
	TechnicBrigadeID    int        `json:"technic_brigade_id"`
	PutIntoService      *time.Time `json:"put_into_service"`
	BaseStationID       *int       `json:"base_station_id"`
}

type LocomotiveFilter struct {
	StationID           *int       // Станция приписки
	ArrivalDate         *time.Time // Дата прибытия на станцию
	EndedRoutesCountMin *int       // Минимальное количество завершённых маршрутов
	EndedRoutesCountMax *int       // Максимальное количество завершённых маршрутов
	Status              *string    // Статус локомотива
	RepairStartDateMin  *time.Time // Минимальная дата начала ремонта
	RepairStartDateMax  *time.Time // Максимальная дата начала ремонта
	RepairEndDateMin    *time.Time // Минимальная дата окончания ремонта
	RepairEndDateMax    *time.Time // Максимальная дата окончания ремонта
	RepairType          *string    // Тип ремонта (например, "плановый")
	RepairCountMin      *int       // Минимальное количество ремонтов
	RepairCountMax      *int       // Максимальное количество ремонтов
	AgeMin              *int       // Минимальный возраст локомотива
	AgeMax              *int       // Максимальный возраст локомотива
}
