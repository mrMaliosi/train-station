package handler

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
)

type EmployeeHandler struct {
	DB           *sqlx.DB
	EmployeeRepo repository.EmployeeRepository
}

// FilterHandler godoc
// @Summary      Получение работников по фильтрам
// @Description  Возвращает список работников, отфильтрованных по возрасту, опыту, полу, детям, зарплате и отделу.
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        department_id  query  int     false  "ID отдела"
// @Param        sex            query  string  false  "Пол (M/F)"
// @Param        age_from       query  int     false  "Возраст от"
// @Param        age_to         query  int     false  "Возраст до"
// @Param        experience_from query int     false  "Опыт от (в годах)"
// @Param        experience_to  query int      false  "Опыт до (в годах)"
// @Param        child_from     query  int     false  "Количество детей от"
// @Param        child_to       query  int     false  "Количество детей до"
// @Param        salary_from    query  number  false  "Зарплата от"
// @Param        salary_to      query  number  false  "Зарплата до"
// @Success      200  {array}   models.Employee
// @Failure      500  {object}  gin.H
// @Router       /employees [get]
func (h *EmployeeHandler) GetFilteredEmployees(c *gin.Context) {
	var filter models.EmployeeFilter

	// Helper: int query param (возвращает nil, если параметр не передан)
	parseInt := func(key string) *int {
		if val := c.DefaultQuery(key, ""); val != "" {
			if v, err := strconv.Atoi(val); err == nil {
				return &v
			}
		}
		return nil
	}

	// Helper: float query param (возвращает nil, если параметр не передан)
	parseFloat := func(key string) *float64 {
		if val := c.DefaultQuery(key, ""); val != "" {
			if v, err := strconv.ParseFloat(val, 64); err == nil {
				return &v
			}
		}
		return nil
	}

	// Чтение параметров
	filter.DepartmentID = parseInt("department_id")
	if sex := c.DefaultQuery("sex", ""); sex != "" {
		filter.Sex = &sex
	}
	filter.AgeFrom = parseInt("age_from")
	filter.AgeTo = parseInt("age_to")
	filter.ExperienceFrom = parseInt("experience_from")
	filter.ExperienceTo = parseInt("experience_to")
	filter.ChildrenFrom = parseInt("child_from")
	filter.ChildrenTo = parseInt("child_to")
	filter.SalaryFrom = parseFloat("salary_from")
	filter.SalaryTo = parseFloat("salary_to")

	// Запрос в репозиторий для получения данных с фильтрацией
	employees, err := h.EmployeeRepo.EmployeeFilter(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ответ с данными
	c.JSON(http.StatusOK, employees)
}
