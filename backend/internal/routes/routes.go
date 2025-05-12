package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/handler"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
	// Инициализация репозиториев
	brigadeRepo := repository.NewBrigadeRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	driverRepo := repository.NewLocomotiveDriverRepository(db)
	locomotiveRepo := repository.NewLocomotiveRepository(db)
	trainRepo := repository.NewTrainRepository(db)
	routeRepo := repository.NewRouteRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	passengerRepo := repository.NewPassengerRepository(db)

	// Инициализация хендлеров
	brigadeHandler := handler.BrigadeHandler{
		DB:          db,
		BrigadeRepo: brigadeRepo,
	}
	employeeHandler := handler.EmployeeHandler{
		DB:           db,
		EmployeeRepo: employeeRepo,
	}
	driverHandler := handler.LocomotiveDriverHandler{
		DriverRepo: driverRepo,
	}
	locomotiveHandler := handler.LocomotiveHandler{
		LocomotiveRepo: locomotiveRepo,
	}
	trainHandler := handler.TrainHandler{
		TrainRepo: trainRepo,
	}
	routeHandler := handler.RouteHandler{
		RouteRepo: routeRepo,
	}
	ticketHandler := handler.TicketHandler{
		TicketRepo: ticketRepo,
	}
	passengerHandler := handler.PassengerHandler{ // Добавлен хэндлер для пассажиров
		PassengerRepo: passengerRepo,
	}

	// Маршруты для бригад
	brigadeGroup := r.Group("/brigades")
	{
		brigadeGroup.GET("/employees", brigadeHandler.GetBrigadeEmployees)
		brigadeGroup.GET("/employees/count", brigadeHandler.CountEmployees)
	}

	// Маршруты для сотрудников
	employeeGroup := r.Group("/employees")
	{
		employeeGroup.GET("", employeeHandler.GetFilteredEmployees)
		employeeGroup.GET("/locomotive-drivers", driverHandler.GetLocomotiveDrivers)
	}

	// Маршруты для локомотивов
	locomotiveGroup := r.Group("/locomotives")
	{
		locomotiveGroup.GET("", locomotiveHandler.GetLocomotives)
		locomotiveGroup.GET("/count", locomotiveHandler.GetLocomotivesCount)
	}

	// Маршруты для поездов
	trainGroup := r.Group("/trains")
	{
		trainGroup.GET("", trainHandler.GetTrains)
		trainGroup.GET("/count", trainHandler.GetTrainsCount)
	}

	// Маршруты для маршрутов
	routeGroup := r.Group("/routes")
	{
		routeGroup.GET("/filter", routeHandler.GetFilteredRoutes)
		routeGroup.GET("/count", routeHandler.CountFilteredRoutes)
		routeGroup.GET("/returned-tickets", routeHandler.GetReturnedTicketsDuringDelay)
	}

	// Маршруты для билетов
	ticketGroup := r.Group("/tickets")
	{
		ticketGroup.GET("", ticketHandler.GetTickets) // Обобщённый эндпоинт для получения билетов с фильтрами
		ticketGroup.GET("/sold/count", ticketHandler.CountSoldTickets)
		ticketGroup.GET("/returned/count", ticketHandler.CountReturnedTickets)
		ticketGroup.GET("/statuses", ticketHandler.GetTicketStatuses) // Получение возможных статусов (ENUM)
		ticketGroup.GET("/stats", ticketHandler.GetTicketStats)

	}

	// Маршруты для пассажиров
	passengerGroup := r.Group("/passengers")
	{
		passengerGroup.GET("/filter", passengerHandler.GetFilteredPassengers) // Добавлен новый маршрут для фильтрации пассажиров
	}
}
