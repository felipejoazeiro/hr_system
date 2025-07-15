package repository

import "app/internal/positions/models"

type PositionRepositoryInterface interface {
	CreatePosition(input models.CreatePosition) (models.PositionModel, error)
	GetAllPositions() ([]models.PositionModel, error)
	EditPosition(id string, input models.EditPosition) (models.PositionModel, error)
	DeactivatePosition(id string) error
	ReactivePosition(id string) error
}
