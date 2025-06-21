package repository

import "time"

func (r *EmployeeRepository) DeactivateEmployeeRepository(employeeID, reason string) error {
	query := `
        UPDATE employees 
        SET 
            active = false,
            deactivation_reason = ?,
            deactivated_at = ?
        WHERE id = ?
    `

	_, err := r.db.Exec(query, reason, time.Now(), employeeID)
	return err
}
