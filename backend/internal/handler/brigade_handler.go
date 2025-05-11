package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

type BrigadeHandler struct {
	DB          *sqlx.DB
	BrigadeRepo repository.BrigadeRepository
}

// Парсинг фильтров
func parseBrigadeEmployeeFilter(c *gin.Context) models.BrigadeEmployeeFilter {
	filter := models.BrigadeEmployeeFilter{}

	if v := c.DefaultQuery("brigade_id", ""); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.BrigadeID = &id
		}
	}
	if v := c.DefaultQuery("department_id", ""); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.DepartmentID = &id
		}
	}
	if v := c.DefaultQuery("locomotive_id", ""); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			filter.LocomotiveID = &id
		}
	}
	if v := c.DefaultQuery("age_from", ""); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			filter.AgeFrom = &val
		}
	}
	if v := c.DefaultQuery("age_to", ""); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			filter.AgeTo = &val
		}
	}
	if v := c.DefaultQuery("salary_from", ""); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			filter.SalaryFrom = &val
		}
	}
	if v := c.DefaultQuery("salary_to", ""); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			filter.SalaryTo = &val
		}
	}

	return filter
}

// Получить сотрудников бригады
func (h *BrigadeHandler) GetBrigadeEmployees(c *gin.Context) {
	filter := parseBrigadeEmployeeFilter(c)

	employees, err := h.BrigadeRepo.GetEmployees(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employees: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// Подсчитать сотрудников
func (h *BrigadeHandler) CountEmployees(c *gin.Context) {
	filter := parseBrigadeEmployeeFilter(c)

	count, err := h.BrigadeRepo.CountEmployees(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting employees: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]int{"count": count})
}
