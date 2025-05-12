package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

// PassengerHandler - структура хэндлера для работы с пассажирами.
type PassengerHandler struct {
	PassengerRepo *repository.PassengerRepository
}

// GetFilteredPassengers - хэндлер для получения списка пассажиров по фильтрам.
func (h *PassengerHandler) GetFilteredPassengers(c *gin.Context) {
	var filter models.PassengerFilter

	filter.RouteID = utilities.IntPtr(c, "routeID")
	filter.Sex = utilities.StringPtr(c.Query("sex"))
	filter.MinAge = utilities.IntPtr(c, "minAge")
	filter.MaxAge = utilities.IntPtr(c, "maxAge")
	filter.HasLuggage = utilities.BoolPtr(c.Query("hasLuggage"))
	filter.Abroad = utilities.BoolPtr(c.Query("abroad"))

	travelDateStr := c.Query("travelDate")
	if travelDateStr != "" {
		filter.TravelDate = &travelDateStr
	}

	passengers, err := h.PassengerRepo.GetFilteredPassengers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passengers)
}
