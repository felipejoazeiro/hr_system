package repository

import (
	"app/internal/hr/models"
	"strings"
)

func (r *EmployeeRepository) ListEmployees(filter models.EmployeeFilter) ([]models.EmployeeDataForList, error) {
	// Query base
	query := `
		SELECT 
			e.id,
			e.register,
			e.name,
			e.email,
			p.name as position_name,
			d.name as department_name,
			e.active
		FROM employees e
		LEFT JOIN positions p ON e.position_id = p.id
		LEFT JOIN departments d ON e.department_id = d.id
		WHERE e.active = true`

	var args []interface{}
	var conditions []string

	if filter.IncludeInactive != nil && *filter.IncludeInactive {
		query = strings.Replace(query, "WHERE e.active = true", "WHERE 1=1", 1)
	}

	if filter.Position != nil {
		conditions = append(conditions, "p.name = ?")
		args = append(args, *filter.Position)
	}

	if filter.Department != nil {
		conditions = append(conditions, "d.name = ?")
		args = append(args, *filter.Department)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.EmployeeDataForList
	for rows.Next() {
		var emp models.EmployeeDataForList
		if err := rows.Scan(
			&emp.ID,
			&emp.Register,
			&emp.Name,
			&emp.Email,
			&emp.Position,
			&emp.Department,
			&emp.Active,
		); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}