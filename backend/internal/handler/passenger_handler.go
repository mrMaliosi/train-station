package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

// PassengerHandler - структура хэндлера для работы с пассажирами.
type PassengerHandler struct {
	PassengerRepo *repository.PassengerRepository
}

// GetFilteredPassengers - хэндлер для получения списка пассажиров по фильтрам.
func (h *PassengerHandler) GetFilteredPassengers(c *gin.Context) {
	var filter models.PassengerFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter parameters"})
		return
	}

	passengers, err := h.PassengerRepo.GetFilteredPassengers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passengers)
}
