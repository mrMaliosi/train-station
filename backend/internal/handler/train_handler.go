package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type TrainHandler struct {
	TrainRepo repository.TrainRepository
}

func NewTrainHandler(repo repository.TrainRepository) *TrainHandler {
	return &TrainHandler{TrainRepo: repo}
}

func (h *TrainHandler) GetTrains(c *gin.Context) {
	routeID := utilities.IntPtr(c, "route_id")
	priceMin := utilities.FloatPtr(c, "price_min")
	priceMax := utilities.FloatPtr(c, "price_max")
	routeTime := utilities.IntPtr(c, "route_time")

	filter := models.TrainFilter{
		RouteID:   routeID,
		PriceMin:  priceMin,
		PriceMax:  priceMax,
		RouteTime: routeTime,
	}

	trains, err := h.TrainRepo.GetTrains(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trains)
}

func (h *TrainHandler) GetTrainsCount(c *gin.Context) {
	routeID := utilities.IntPtr(c, "route_id")
	priceMin := utilities.FloatPtr(c, "price_min")
	priceMax := utilities.FloatPtr(c, "price_max")
	routeTime := utilities.IntPtr(c, "route_time")

	filter := models.TrainFilter{
		RouteID:   routeID,
		PriceMin:  priceMin,
		PriceMax:  priceMax,
		RouteTime: routeTime,
	}

	count, err := h.TrainRepo.GetTrainsCount(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
