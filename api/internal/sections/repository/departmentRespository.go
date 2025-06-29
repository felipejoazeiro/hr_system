package repository

import (
	"app/internal/sections/models"
	"database/sql"
	"fmt"
	"strings"
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

func (r *DepartmentRepository) EditDepartment(id string, d models.EditDepartmentRequest) (models.DepartmentModel, error) {
    setParts := []string{}
    args := []interface{}{}
    argPos := 1

    if d.Name != nil {
        setParts = append(setParts, fmt.Sprintf("name=$%d", argPos))
        args = append(args, *d.Name)
        argPos++
    }
    if d.Code != nil {
        setParts = append(setParts, fmt.Sprintf("code=$%d", argPos))
        args = append(args, *d.Code)
        argPos++
    }
    if d.ManagerId != nil {
        setParts = append(setParts, fmt.Sprintf("manager_id=$%d", argPos))
        args = append(args, *d.ManagerId)
        argPos++
    }

    if len(setParts) == 0 {
        return models.DepartmentModel{}, nil
    }

    args = append(args, id)
    query := fmt.Sprintf("UPDATE departments SET %s WHERE id=$%d RETURNING id, name, code, manager_id, is_active", 
        strings.Join(setParts, ", "), argPos)

    var updated models.DepartmentModel
    err := r.db.QueryRow(query, args...).Scan(
        &updated.ID,
        &updated.Name,
        &updated.Code,
        &updated.ManagerId,
        &updated.IsActive,
    )
    return updated, err
}

func (r *DepartmentRepository) DeactivateDepartment(id string) error {
	query := `UPDATE departments SET is_active = false WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *DepartmentRepository) ReactivateDepartment(id string) error {
	query := `UPDATE departments SET is_active = true WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
