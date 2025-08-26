package repository

import (
	"database/sql"
)

type FinancePayrollRepository struct {
	db *sql.DB
}

func NewFinancePayrollRepository(db *sql.DB) *FinancePayrollRepository {
	return &FinancePayrollRepository{
		db: db,
	}
}

func (r *FinancePayrollRepository) GetAllPayrollComponents() ([]models.PayrollComponents, error) {
	var components []models.PayrollComponents
	err := r.db.Select(&components, "SELECT * FROM payroll_components")
	if err != nil {
		return nil, err
	}
	return components, nil
}

func (r *FinancePayrollRepository) GetPayrollComponentByID(id int) (*models.PayrollComponents, error) {
	var component models.PayrollComponents
	err := r.db.Get(&component, "SELECT * FROM payroll_components WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &component, nil
}

func (r *FinancePayrollRepository) CreatePayrollComponent(component models.PayrollComponents) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO payroll_components (field1, field2) VALUES (?, ?) RETURNING id",
		component.Field1, component.Field2).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FinancePayrollRepository) UpdatePayrollComponent(input models.PayrollComponents) error {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if input.Code != "" {
		setParts = append(setParts, fmt.Sprintf("code = $%d", argPos))
		args = append(args, input.Code)
		argPos++
	}
	if input.Description != "" {
		setParts = append(setParts, fmt.Sprintf("description = $%d", argPos))
		args = append(args, input.Description)
		argPos++
	}
	if input.Nature != "" {
		setParts = append(setParts, fmt.Sprintf("nature = $%d", argPos))
		args = append(args, input.Nature)
		argPos++
	}
	if input.FkAccountID != 0 {
		setParts = append(setParts, fmt.Sprintf("fk_accountId = $%d", argPos))
		args = append(args, input.FkAccountID)
		argPos++
	}

	args = append(args, input.ID)

	query := fmt.Sprintf("UPDATE payroll_components SET %s WHERE id = ?", strings.Join(setParts, ", "))
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *FinancePayrollRepository) DeactivatePayrollComponent(id int) error {
	_, err := r.db.Exec("UPDATE payroll_components SET active = false WHERE id = ?", id)
	return err
}

func (r *FinancePayrollRepository) ReactivatePayrollComponent(id int) error {
	_, err := r.db.Exec("UPDATE payroll_components SET active = true WHERE id = ?", id)
	return err
}

func (r *FinancePayrollRepository) GetAllEmployeeSalaryHistories() ([]models.EmployeeSalaryHistory, error) {
	var histories []models.EmployeeSalaryHistory
	err := r.db.Select(&histories, "SELECT * FROM employee_salary_history")
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *FinancePayrollRepository) GetEmployeeSalaryHistoriesById(id int) (*models.EmployeeSalaryHistory, error) {
	var history models.EmployeeSalaryHistory
	err := r.db.Get(&history, "SELECT * FROM employee_salary_history WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *FinancePayrollRepository) CreateEmployeeSalaryHistory(history models.CreateEmployeeSalaryHistory) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO employee_salary_history (field1, field2) VALUES (?, ?) RETURNING id",
		history.Field1, history.Field2).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FinancePayrollRepository) EditEmployeeSalaryHistory(input models.EditEmployeeSalaryHistory) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if input.BaseSalary != nil {
		setParts = append(setParts, fmt.Sprintf("base_salary = $%d", argPos))
		args = append(args, *input.BaseSalary)
		argPos++
	}

	if input.ValidFrom != nil {
		setParts = append(setParts, fmt.Sprintf("valid_from = $%d", argPos))
		args = append(args, *input.ValidFrom)
		argPos++
	}

	if input.ValidTo != nil {
		setParts = append(setParts, fmt.Sprintf("valid_to = $%d", argPos))
		args = append(args, *input.ValidTo)
		argPos++
	}

	if input.FkEmployeeID != 0 {
		setParts = append(setParts, fmt.Sprintf("fk_employee_id = $%d", argPos))
		args = append(args, input.FkEmployeeID)
		argPos++
	}

	args = append(args, input.ID)

	query := fmt.Sprintf("UPDATE employee_salary_history SET %s WHERE id = ?", strings.Join(setParts, ", "))
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *FinancePayrollRepository) GetAllPayrollRuns() ([]models.PayrollRuns, error) {
	var runs []models.PayrollRuns
	err := r.db.Select(&runs, "SELECT * FROM payroll_runs")
	if err != nil {
		return nil, err
	}
	return runs, nil
}

func (r *FinancePayrollRepository) GetPayrollRunByID(id int) (*models.PayrollRuns, error) {
	var run models.PayrollRuns
	err := r.db.Get(&run, "SELECT * FROM payroll_runs WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &run, nil
}

func (r *FinancePayrollRepository) CreatePayrollRun(input models.CreatePayrollRuns) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if input.PeriodStart != nil {
		setParts = append(setParts, fmt.Sprintf("period_start = $%d", argPos))
		args = append(args, *input.PeriodStart)
		argPos++
	}

	if input.PeriodEnd != nil {
		setParts = append(setParts, fmt.Sprintf("period_end = $%d", argPos))
		args = append(args, *input.PeriodEnd)
		argPos++
	}

	if input.Status != "" {
		setParts = append(setParts, fmt.Sprintf("status = $%d", argPos))
		args = append(args, input.Status)
		argPos++
	}

	if input.Notes != nil {
		setParts = append(setParts, fmt.Sprintf("notes = $%d", argPos))
		args = append(args, *input.Notes)
		argPos++
	}

	query := fmt.Sprintf("INSERT INTO payroll_runs SET %s RETURNING id", strings.Join(setParts, ", "))
	err := r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FinancePayrollRepository) GetAllPayrollItems() ([]models.PayrollItems, error) {
	var items []models.PayrollItems
	err := r.db.Select(&items, "SELECT * FROM payroll_items")
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *FinancePayrollRepository) GetPayrollItemByID(id int) (*models.PayrollItems, error) {
	var item models.PayrollItems
	err := r.db.Get(&item, "SELECT * FROM payroll_items WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *FinancePayrollRepository) CreatePayrollItem(input models.CreatePayrollItems) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO payroll_items (field1, field2) VALUES (?, ?) RETURNING id",
		input.Field1, input.Field2).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FinancePayrollRepository) EditPayrollItem(input models.EditPayrollItems) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if input.Amount != nil {
		setParts = append(setParts, fmt.Sprintf("amount = $%d", argPos))
		args = append(args, *input.Amount)
		argPos++
	}

	if input.FkPayrollRunId != 0 {
		setParts = append(setParts, fmt.Sprintf("fk_payroll_run_id = $%d", argPos))
		args = append(args, input.FkPayrollRunId)
		argPos++
	}

	if input.FkEmployeeId != 0 {
		setParts = append(setParts, fmt.Sprintf("fk_employee_id = $%d", argPos))
		args = append(args, input.FkEmployeeId)
		argPos++
	}

	if input.FkComponentId != 0 {
		setParts = append(setParts, fmt.Sprintf("fk_component_id = $%d", argPos))
		args = append(args, input.FkComponentId)
		argPos++
	}

	if input.FkCostCenterId != 0 {
		setParts = append(setParts, fmt.Sprintf("fk_cost_center_id = $%d", argPos))
		args = append(args, input.FkCostCenterId)
		argPos++
	}

	args = append(args, input.ID)

	query := fmt.Sprintf("UPDATE payroll_items SET %s WHERE id = ?", strings.Join(setParts, ", "))
	_, err := r.db.Exec(query, args...)
	return err
}