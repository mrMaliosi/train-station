package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type LocomotiveDriverHandler struct {
	DriverRepo repository.LocomotiveDriverRepository
}

// GetLocomotiveDrivers godoc
// @Summary      Получение водителей локомотивов по фильтрам
// @Description  Возвращает список водителей локомотивов, прошедших или не прошедших медосмотр в указанном году
// @Tags         locomotive_drivers
// @Accept       json
// @Produce      json
// @Param        sex             query  string  false  "Пол (M/F)"
// @Param        age_from        query  int     false  "Возраст от"
// @Param        age_to          query  int     false  "Возраст до"
// @Param        salary_from     query  number  false  "Зарплата от"
// @Param        salary_to       query  number  false  "Зарплата до"
// @Param        medical_checkup_year query int false "Год медосмотра"
// @Success      200  {array}   models.Employee
// @Failure      500  {object}  gin.H
// @Router       /locomotive-drivers [get]
func (h *LocomotiveDriverHandler) GetLocomotiveDrivers(c *gin.Context) {
	var filter models.LocomotiveDriverFilter

	// Чтение параметров запроса
	filter.Sex = utilities.StringPtr(c.DefaultQuery("sex", ""))
	filter.AgeFrom = utilities.IntPtr(c, "age_from")
	filter.AgeTo = utilities.IntPtr(c, "age_to")
	filter.SalaryFrom = utilities.FloatPtr(c, "salary_from")
	filter.SalaryTo = utilities.FloatPtr(c, "salary_to")
	filter.MedicalCheckupYear = utilities.IntPtr(c, "medical_checkup_year")

	// Получение списка водителей
	drivers, err := h.DriverRepo.GetLocomotiveDrivers(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Получение общего числа водителей
	count, err := h.DriverRepo.GetLocomotiveDriversCount(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"drivers": drivers,
		"count":   count,
	})
}
