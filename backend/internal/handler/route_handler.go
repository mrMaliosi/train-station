package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type RouteHandler struct {
	RouteRepo repository.RouteRepository
}

func (h *RouteHandler) GetFilteredRoutes(c *gin.Context) {
	filter := models.RouteFilter{
		RouteID:     utilities.IntPtr(c, "route_id"),
		Status:      utilities.StringPtr(c.Query("status")),
		DelayReason: utilities.StringPtr(c.Query("reason")),
		TrainType:   utilities.StringPtr(c.Query("train_type")),
		StationName: utilities.StringPtr(c.Query("station_name")),
	}

	routes, err := h.RouteRepo.GetFilteredRoutes(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, routes)
}

func (h *RouteHandler) CountFilteredRoutes(c *gin.Context) {
	filter := models.RouteFilter{
		RouteID:     utilities.IntPtr(c, "route_id"),
		Status:      utilities.StringPtr(c.Query("status")),
		DelayReason: utilities.StringPtr(c.Query("reason")),
		TrainType:   utilities.StringPtr(c.Query("train_type")),   // добавлено
		StationName: utilities.StringPtr(c.Query("station_name")), // добавлено
	}

	count, err := h.RouteRepo.CountFilteredRoutes(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *RouteHandler) GetReturnedTicketsDuringDelay(c *gin.Context) {
	filter := models.RouteFilter{
		RouteID:     utilities.IntPtr(c, "route_id"),
		DelayReason: utilities.StringPtr(c.Query("reason")),
	}

	count, err := h.RouteRepo.GetReturnedTicketsDuringDelay(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"returned": count})
}
