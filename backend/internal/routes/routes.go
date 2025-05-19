package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/handler"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173" || origin == "http://127.0.0.1:5173"
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Инициализация репозиториев
	brigadeRepo := repository.NewBrigadeRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	driverRepo := repository.NewLocomotiveDriverRepository(db)
	locomotiveRepo := repository.NewLocomotiveRepository(db)
	trainRepo := repository.NewTrainRepository(db)
	routeRepo := repository.NewRouteRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	passengerRepo := repository.NewPassengerRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)
	positionRepo := repository.NewPositionRepository(db)

	// Инициализация хендлеров
	brigadeHandler := handler.BrigadeHandler{
		DB:          db,
		BrigadeRepo: brigadeRepo,
	}
	employeeHandler := handler.EmployeeHandler{
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
	passengerHandler := handler.PassengerHandler{
		PassengerRepo: passengerRepo,
	}
	departmentHandler := handler.DepartmentHandler{
		DepartmentRepo: departmentRepo,
	}
	positionHandler := handler.PositionHandler{
		PositionRepo: positionRepo,
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
		employeeGroup.POST("", employeeHandler.PostNewEmployee)
		employeeGroup.DELETE("/:id", employeeHandler.DeleteEmployee)
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
		passengerGroup.GET("/filter", passengerHandler.GetFilteredPassengers)
	}

	// Маршруты для департаментов
	departmentGroup := r.Group("/departments")
	{
		departmentGroup.GET("", departmentHandler.GetDepartments)
		departmentGroup.GET("/info", departmentHandler.GetDepartmentsInfo)
	}

	// Маршруты для позиций
	positionGroup := r.Group("/positions")
	{
		positionGroup.GET("", positionHandler.GetPositions)
	}
}
