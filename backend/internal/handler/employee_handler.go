package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrMaliosi/train-station/backend/internal/models"
	"github.com/mrMaliosi/train-station/backend/internal/repository"
	"github.com/mrMaliosi/train-station/backend/internal/utilities"
)

type EmployeeHandler struct {
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

	// Чтение параметров
	filter.DepartmentID = utilities.IntPtr(c, "department_id")
	if sex := c.DefaultQuery("sex", ""); sex != "" {
		filter.Sex = &sex
	}
	filter.AgeFrom = utilities.IntPtr(c, "age_from")
	filter.AgeTo = utilities.IntPtr(c, "age_to")
	filter.ExperienceFrom = utilities.IntPtr(c, "experience_from")
	filter.ExperienceTo = utilities.IntPtr(c, "experience_to")
	filter.ChildrenFrom = utilities.IntPtr(c, "children_from")
	filter.ChildrenTo = utilities.IntPtr(c, "children_to")
	filter.SalaryFrom = utilities.FloatPtr(c, "salary_from")
	filter.SalaryTo = utilities.FloatPtr(c, "salary_to")

	// Запрос в репозиторий для получения данных с фильтрацией
	employees, err := h.EmployeeRepo.EmployeeFilter(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ответ с данными
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) PostNewEmployee(c *gin.Context) {
	var req models.EmployeeCreate

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employeeID, err := h.EmployeeRepo.EmployeeCreate(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employeeID)
}
