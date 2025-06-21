package repository

import (
	"database/sql"
	"fmt"
	"app/internal/hr/models"
)

// Variável global do repositório (ou injete via DI)
var Repository *EmployeeRepository

type EmployeeRepository struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) {
	Repository = NewEmployeeRepository(db)
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) GetUserWithPosition(login string) (*models.EmployeeWithPosition, error) {
	query := `
		SELECT e.login, e.password, c.name as position_name
		FROM employee_access e
		JOIN employee emp ON e.employee_id = emp.id
		JOIN cargo c ON emp.fk_cargo = c.id
		WHERE e.login = ?
	`
	
	row := r.db.QueryRow(query, login)
	
	var employee models.EmployeeWithPosition
	err := row.Scan(&employee.Login, &employee.Password, &employee.Position)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	
	return &employee, nil
}