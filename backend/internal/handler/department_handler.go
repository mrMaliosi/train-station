package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

type DepartmentHandler struct {
	DepartmentRepo repository.DepartmentRepository
}

func (h *DepartmentHandler) GetDepartments(c *gin.Context) {
	// Запрос в репозиторий для получения данных с фильтрацией
	departments, err := h.DepartmentRepo.DepartmentsSelect(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ответ с данными
	c.JSON(http.StatusOK, departments)
}
