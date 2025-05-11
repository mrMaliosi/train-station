package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type TicketHandler struct {
	TicketRepo repository.TicketRepository
}

func (h *TicketHandler) GetSoldTickets(c *gin.Context) {
	filter := models.TicketFilter{
		FromDate: utilities.StringPtr(c.Query("from")),
		ToDate:   utilities.StringPtr(c.Query("to")),
		RouteID:  utilities.IntPtr(c, "route_id"),
	}
	tickets, err := h.TicketRepo.GetSoldTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) CountSoldTickets(c *gin.Context) {
	filter := models.TicketFilter{
		FromDate: utilities.StringPtr(c.Query("from")),
		ToDate:   utilities.StringPtr(c.Query("to")),
		RouteID:  utilities.IntPtr(c, "route_id"),
	}
	count, err := h.TicketRepo.CountSoldTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (h *TicketHandler) GetUnsoldTickets(c *gin.Context) {
	filter := models.TicketFilter{
		Date:    utilities.StringPtr(c.Query("date")),
		RouteID: utilities.IntPtr(c, "route_id"),
	}
	tickets, err := h.TicketRepo.GetUnsoldTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) CountReturnedTickets(c *gin.Context) {
	filter := models.TicketFilter{
		Date:    utilities.StringPtr(c.Query("date")),
		RouteID: utilities.IntPtr(c, "route_id"),
	}
	count, err := h.TicketRepo.CountReturnedTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"returned": count})
}
