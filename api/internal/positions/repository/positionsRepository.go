package repository

import (
	"app/internal/positions/models"
	"database/sql"
	"fmt"
	"strings"
)

type PositionRepository struct {
	db *sql.DB
}

func NewPositionRepository(db *sql.DB) *PositionRepository {
	return &PositionRepository{db: db}
}

func (r *PositionRepository) CreatePosition(input models.CreatePosition) (models.PositionModel, error) {
	query := `
		INSERT INTO positions (name, code, is_active)
		VALUES ($1, $2, true)
		RETURNING id, name, code, is_active
	`

	var pos models.PositionModel
	err := r.db.QueryRow(
		query,
		input.Name,
		input.Code,
	).Scan(
		&pos.ID,
		&pos.Name,
		&pos.Code,
		&pos.IsActive,
	)
	return pos, err
}

func (r *PositionRepository) GetAllPositions() ([]models.PositionModel, error) {
	query := `SELECT id, name, code, is_active FROM positions ORDER BY id;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []models.PositionModel
	for rows.Next() {
		var pos models.PositionModel
		if err := rows.Scan(
			&pos.ID,
			&pos.Name,
			&pos.Code,
			&pos.IsActive,
		); err != nil {
			return nil, err
		}
		positions = append(positions, pos)
	}
	return positions, nil
}

func (r *PositionRepository) EditPosition(id string, input models.EditPosition) (models.PositionModel, error) {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if input.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, *input.Name)
		argIndex++
	}
	if input.Code != nil {
		setParts = append(setParts, fmt.Sprintf("code = $%d", argIndex))
		args = append(args, *input.Code)
		argIndex++
	}

	if len(setParts) == 0 {
		return models.PositionModel{}, fmt.Errorf("no fields to update")
	}

	// ID entra como último argumento
	query := fmt.Sprintf(
		"UPDATE positions SET %s WHERE id = $%d RETURNING id, name, code, is_active",
		strings.Join(setParts, ", "),
		argIndex,
	)
	args = append(args, id)

	var pos models.PositionModel
	err := r.db.QueryRow(query, args...).Scan(
		&pos.ID,
		&pos.Name,
		&pos.Code,
		&pos.IsActive,
	)
	return pos, err
}

func (r *PositionRepository) DeactivatePosition(id string) error {
	_, err := r.db.Exec(`UPDATE positions SET is_active = false WHERE id = $1`, id)
	return err
}

func (r *PositionRepository) ReactivePosition(id string) error {
	_, err := r.db.Exec(`UPDATE positions SET is_active = true WHERE id = $1`, id)
	return err
}
