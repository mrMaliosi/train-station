package utilities

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// IntPtr извлекает параметр запроса и возвращает *int (nil, если параметр не передан или невалиден)
func IntPtr(c *gin.Context, key string) *int {
	if val := c.DefaultQuery(key, ""); val != "" {
		if v, err := strconv.Atoi(val); err == nil {
			return &v
		}
	}
	return nil
}

// FloatPtr извлекает параметр запроса и возвращает *float64 (nil, если параметр не передан или невалиден)
func FloatPtr(c *gin.Context, key string) *float64 {
	if val := c.DefaultQuery(key, ""); val != "" {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			return &v
		}
	}
	return nil
}

// StringPtr возвращает *string, если строка не пустая; иначе nil
func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// ParseTime пытается распарсить строку даты в формат "2006-01-02" и вернуть указатель на time.Time.
// Если строка пуста или формат даты неверный, возвращает nil.
func ParseTime(c *gin.Context, queryParam string) *time.Time {
	dateStr := c.Query(queryParam)
	if dateStr == "" {
		return nil
	}
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// Логируем ошибку, если дата невалидна
		log.Printf("Invalid date format for %s: %s, error: %v", queryParam, dateStr, err)
		return nil
	}
	return &parsedDate
}
