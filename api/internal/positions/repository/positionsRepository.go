package repository



type PositionRepository struct {
	db *sql.DB
}

func NewPositionRepository(db *sql.DB) *PositionRepository{
	return &PositionRepository{db: db}
}

func (r *PositionRepository) CreatePosition(input models.CreatePosition) (models.PositionModel, error){
	query := `
		INSERT INTO positons (name, code, is_active)
		VALUES ($1, $2, true)
		RETURNING id, name, code, is_active
	`

	var pos models.PositionModel
	err := r.db.QueryRow(
		query, 
		input.Name,
		input.Code
	).Scan(
		&pos.ID,
		&pos.Name,
		&pos.Code,
		&pos.IsActive,
	)
	return pos, err
}	

func (r *PositionRepository) GetAllPositions() ([]models.PositionModel, error) {
	query := `SELECT id, name, code, is_active FROM positions ORDER BY id;`;

	rows, err := r.db.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var positions []models.PositionModel
	for rows.Next(){
		var pos models.PositionModel
		if err := rows.Scan(
			&pos.ID,
			&pos.Name,
			&pos.Code,
			&pos.IsActive
		); err := nil {
			return nil, err
		}
		positions = append(positions, pos)
	}
	return positions,nil
}

func (r *PositionRepository) EditPosition(id string, input models.EditPosition) (models.PositionModel, error){
	setParts := []string{}
	args := []interface{}{}

	if input.Name != nil{
		setParts = append(setParts, "name = ?")
		args = append(args, *input.Name)
	}
	if input.Code != nil {
		setParts = append(setParts, "code = ?")
		args = append(args, *input.Code)
	}
	if input.IsActive != nil {
		setParts = append(setParts, "is_active = ?")
		args = append(args, *input.IsActive)
	}

	if len(setParts) == 0 {
		return models.PositionModel{}, nil
	}

	query := "UPDATE positions SET " + strings.Join(setParts, ", ") + " WHERE id = ? RETURNING id, name, code, is_active"
	args = append(args, id)

	var pos models.PositionModel
	err := r.db.QueryRow(
		query,
		args...,
	).Scan(
		&pos.ID,
		&pos.Name,
		&pos.Code,
		&pos.IsActive,
	)
	return pos, err
}

func (r *PositionRepository) DeactivatePosition(id string) (models.PositionModel, error){}

func (r *PositionRepository) ReactivePosition(id string) (models.PositionModel, error){}