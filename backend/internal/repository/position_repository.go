package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrMaliosi/train-station/backend/internal/models"
)

// Интерфейс для репозитория сотрудников
type PositionRepository interface {
	PositionsSelect(ctx context.Context) ([]models.PositionName, error)
}

// Реализация репозитория сотрудников
type positionRepository struct {
	db *sqlx.DB
}

// Конструктор для репозитория сотрудников
func NewPositionRepository(db *sqlx.DB) PositionRepository {
	return &positionRepository{db: db}
}

func (r *positionRepository) PositionsSelect(
	ctx context.Context,
) ([]models.PositionName, error) {
	baseQuery := `
    SELECT d.position_id, d.position_name
	FROM "Positions" AS d`

	fmt.Println("Executing query:", baseQuery)

	var positions []models.PositionName
	err := r.db.SelectContext(ctx, &positions, baseQuery)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	return positions, nil
}
