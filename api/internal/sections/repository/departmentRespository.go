package repository

import (
	"app/internal/sections/models"
	"database/sql"
)

type DepartmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) GetAllDepartments() ([]models.GetDepartmentsRequest, error) {
	query := `SELECT d.id, d.name, d.code, e.manager_name, d.is_active FROM department d LEFT JOIN employees e ON d.manager_id = e.id;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []models.GetDepartmentsRequest

	for rows.Next() {
		var dept models.GetDepartmentsRequest
		err := rows.Scan(
			&dept.ID,
			&dept.Name,
			&dept.Code,
			&dept.ManagerName,
		)
		if err != nil {
			return nil, err
		}
		departments = append(departments, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *DepartmentRepository) CreateDepartment(d models.CreateDepartmentRequest) (models.DepartmentModel, error) {
	query := `INSERT INTO departments (name, code, manager_id, is_active)
		VALUES ($1, $2, $3, TRUE)
		RETURNING id, name, code, manager_id, is_active;`

	var department models.DepartmentModel

	err := r.db.QueryRow(query, d.Name, d.Code, d.ManagerId).Scan(
		&department.ID,
		&department.Name,
		&department.Code,
		&department.ManagerId,
		&department.IsActive,
	)
	return department, err
}
