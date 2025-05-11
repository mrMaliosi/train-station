package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type LocomotiveHandler struct {
	LocomotiveRepo repository.LocomotiveRepository
}

func NewLocomotiveHandler(repo repository.LocomotiveRepository) *LocomotiveHandler {
	return &LocomotiveHandler{LocomotiveRepo: repo}
}

func (h *LocomotiveHandler) GetLocomotives(c *gin.Context) {
	// Получаем параметры запроса
	stationID := utilities.IntPtr(c, "station_id")
	arrivalStr := c.Query("arrival_date")
	var arrivalDate *time.Time
	if arrivalStr != "" {
		if t, err := time.Parse("2006-01-02", arrivalStr); err == nil {
			arrivalDate = &t
		}
	}
	endedMin := utilities.IntPtr(c, "ended_min")
	endedMax := utilities.IntPtr(c, "ended_max")
	status := utilities.StringPtr(c.Query("status"))
	repairStartMin := utilities.ParseTime(c, "repair_start_min")
	repairStartMax := utilities.ParseTime(c, "repair_start_max")
	repairEndMin := utilities.ParseTime(c, "repair_end_min")
	repairEndMax := utilities.ParseTime(c, "repair_end_max")
	repairType := utilities.StringPtr(c.Query("repair_type"))
	repairCountMin := utilities.IntPtr(c, "repair_count_min")
	repairCountMax := utilities.IntPtr(c, "repair_count_max")
	ageMin := utilities.IntPtr(c, "age_min")
	ageMax := utilities.IntPtr(c, "age_max")

	filter := models.LocomotiveFilter{
		StationID:           stationID,
		ArrivalDate:         arrivalDate,
		EndedRoutesCountMin: endedMin,
		EndedRoutesCountMax: endedMax,
		Status:              status,
		RepairStartDateMin:  repairStartMin,
		RepairStartDateMax:  repairStartMax,
		RepairEndDateMin:    repairEndMin,
		RepairEndDateMax:    repairEndMax,
		RepairType:          repairType,
		RepairCountMin:      repairCountMin,
		RepairCountMax:      repairCountMax,
		AgeMin:              ageMin,
		AgeMax:              ageMax,
	}

	locomotives, err := h.LocomotiveRepo.GetLocomotives(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, locomotives)
}

func (h *LocomotiveHandler) GetLocomotivesCount(c *gin.Context) {
	// Получаем параметры запроса
	stationID := utilities.IntPtr(c, "station_id")
	arrivalStr := c.Query("arrival_date")
	var arrivalDate *time.Time
	if arrivalStr != "" {
		if t, err := time.Parse("2006-01-02", arrivalStr); err == nil {
			arrivalDate = &t
		}
	}
	endedMin := utilities.IntPtr(c, "ended_min")
	endedMax := utilities.IntPtr(c, "ended_max")
	status := utilities.StringPtr(c.Query("status"))
	repairStartMin := utilities.ParseTime(c, "repair_start_min")
	repairStartMax := utilities.ParseTime(c, "repair_start_max")
	repairEndMin := utilities.ParseTime(c, "repair_end_min")
	repairEndMax := utilities.ParseTime(c, "repair_end_max")
	repairType := utilities.StringPtr(c.Query("repair_type"))
	repairCountMin := utilities.IntPtr(c, "repair_count_min")
	repairCountMax := utilities.IntPtr(c, "repair_count_max")
	ageMin := utilities.IntPtr(c, "age_min")
	ageMax := utilities.IntPtr(c, "age_max")

	filter := models.LocomotiveFilter{
		StationID:           stationID,
		ArrivalDate:         arrivalDate,
		EndedRoutesCountMin: endedMin,
		EndedRoutesCountMax: endedMax,
		Status:              status,
		RepairStartDateMin:  repairStartMin,
		RepairStartDateMax:  repairStartMax,
		RepairEndDateMin:    repairEndMin,
		RepairEndDateMax:    repairEndMax,
		RepairType:          repairType,
		RepairCountMin:      repairCountMin,
		RepairCountMax:      repairCountMax,
		AgeMin:              ageMin,
		AgeMax:              ageMax,
	}

	count, err := h.LocomotiveRepo.GetLocomotivesCount(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}
