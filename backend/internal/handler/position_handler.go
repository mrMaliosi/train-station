package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

type PositionHandler struct {
	PositionRepo repository.PositionRepository
}

func (h *PositionHandler) GetPositions(c *gin.Context) {
	// Запрос в репозиторий для получения данных с фильтрацией
	departments, err := h.PositionRepo.PositionsSelect(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ответ с данными
	c.JSON(http.StatusOK, departments)
}
