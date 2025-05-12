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

// Универсальный метод для получения билетов с фильтрацией
func (h *TicketHandler) GetTickets(c *gin.Context) {
	filter := models.TicketFilter{
		FromDate: utilities.StringPtr(c.Query("fromDate")),
		ToDate:   utilities.StringPtr(c.Query("toDate")),
		RouteID:  utilities.IntPtr(c, "RouteID"),
		Status:   utilities.StringPtr(c.Query("status")),
	}

	tickets, err := h.TicketRepo.GetTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

// Подсчёт проданных билетов
func (h *TicketHandler) CountSoldTickets(c *gin.Context) {
	filter := models.TicketFilter{
		FromDate: utilities.StringPtr(c.Query("fromDate")),
		ToDate:   utilities.StringPtr(c.Query("toDate")),
		RouteID:  utilities.IntPtr(c, "routeID"),
		Status:   utilities.StringPtr(c.Query("status")),
	}
	count, err := h.TicketRepo.CountSoldTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// Подсчёт возвращённых билетов
func (h *TicketHandler) CountReturnedTickets(c *gin.Context) {
	filter := models.TicketFilter{
		FromDate: utilities.StringPtr(c.Query("fromDate")),
		ToDate:   utilities.StringPtr(c.Query("toDate")),
		RouteID:  utilities.IntPtr(c, "routeID"),
		Status:   utilities.StringPtr(c.Query("status")),
	}
	count, err := h.TicketRepo.CountReturnedTickets(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"returned": count})
}

// Получение доступных статусов из ENUM
func (h *TicketHandler) GetTicketStatuses(c *gin.Context) {
	statuses, err := h.TicketRepo.GetTicketStatuses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot load statuses"})
		return
	}
	c.JSON(http.StatusOK, statuses)
}

func (h *TicketHandler) GetTicketStats(c *gin.Context) {
	// Получаем фильтры
	filter := models.TicketFilter{
		FromDate: utilities.StringPtr(c.Query("fromDate")),
		ToDate:   utilities.StringPtr(c.Query("toDate")),
		RouteID:  utilities.IntPtr(c, "routeID"),
		Status:   utilities.StringPtr(c.Query("status")),
	}
	if err := c.BindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем статистику
	soldCount, err := h.TicketRepo.CountSoldTickets(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при подсчёте проданных билетов"})
		return
	}

	returnedCount, err := h.TicketRepo.CountReturnedTickets(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при подсчёте возвращённых билетов"})
		return
	}

	// Отправляем ответ
	c.JSON(http.StatusOK, gin.H{
		"sold":     soldCount,
		"returned": returnedCount,
	})
}
