package models

import "time"

type PayrollComponents struct {
	ID				int 		`json:"id" db:"id"`
	Code 			string 		`json:"code" db:"code"`
	Name 			string 		`json:"name" db:"name"`
	Nature 			string 		`json:"nature" db:"nature"`
	FkAccountID		int 		`json:"fk_account_id" db:"fk_account_id"`	
	isActive		bool 		`json:"is_active" db:"is_active"`
}

type CreatePayrollComponents struct {
	Code 		string 		`json:"code" db:"code"`
	Name 		string 		`json:"name" db:"name"`
	Nature 		string 		`json:"nature" db:"nature"`
	FkAccountID	int 		`json:"fk_account_id" db:"fk_account_id"`
	isActive	bool 		`json:"is_active" db:"is_active"`
}

type EditPayrollComponents struct {
	Code 		*string 	`json:"code" db:"code"`
	Name 		*string 	`json:"name" db:"name"`
	Nature 		*string 	`json:"nature" db:"nature"`
	FkAccountID	*int 		`json:"fk_account_id" db:"fk_account_id"`
}

type EmployeeSalaryHistory struct {
	ID				int 		`json:"id" db:"id"`
	BaseSalary		float64		`json:"base_salary" db:"base_salary"`
	ValidFrom		time.Time	`json:"valid_from" db:"valid_from"`
	ValidTo			time.Time	`json:"valid_to" db:"valid_to"`
	FkEmployee		int 		`json:"fk_employee" db:"fk_employee"`
}

type CreateEmployeeSalaryHistory struct {
	BaseSalary		float64		`json:"base_salary" db:"base_salary"`
	ValidFrom		time.Time	`json:"valid_from" db:"valid_from"`
	ValidTo			time.Time	`json:"valid_to" db:"valid_to"`
	FkEmployee		int 		`json:"fk_employee" db:"fk_employee"`
}

type EditEmployeeSalaryHistory struct {
	BaseSalary		*float64		`json:"base_salary" db:"base_salary"`
	ValidFrom		*time.Time		`json:"valid_from" db:"valid_from"`
	ValidTo			*time.Time		`json:"valid_to" db:"valid_to"`
	FkEmployee		*int 			`json:"fk_employee" db:"fk_employee"`
}

type PayrollRuns struct {
	ID			int		`json:"id" db:"id"`
	PeriodStart	time.Time	`json:"period_start" db:"period_start"`
	PeriodEnd	time.Time	`json:"period_end" db:"period_end"`
	Status		string		`json:"status" db:"status"`
	Notes		string		`json:"notes" db:"notes"`
}

type CreatePayrollRuns struct {
	PeriodStart	time.Time	`json:"period_start" db:"period_start"`
	PeriodEnd	time.Time	`json:"period_end" db:"period_end"`
	Status		string		`json:"status" db:"status"`
	Notes		string		`json:"notes" db:"notes"`
}

type EditPayrollRuns struct {
	PeriodStart		*time.Time	`json:"period_start" db:"period_start"`
	PeriodEnd		*time.Time	`json:"period_end" db:"period_end"`
	Status			*string		`json:"status" db:"status"`
	Notes			*string		`json:"notes" db:"notes"`
}

type PayrollItems struct {
	ID				int 		`json:"id" db:"id"`
	Amount			float64		`json:"amount" db:"amount"`
	FkPayrollRunId	int 		`json:"fk_payrollRunId" db:"fk_payrollRunId"`
	FkEmployeeId	int 		`json:"fk_employeeId" db:"fk_employeeId"`
	FkComponentId	int 		`json:"fk_componentId" db:"fk_componentId"`
	FkCostCenterId	int 		`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type CreatePayrollItems struct {
	Amount			float64		`json:"amount" db:"amount"`
	FkPayrollRunId	int 		`json:"fk_payrollRunId" db:"fk_payrollRunId"`
	FkEmployeeId	int 		`json:"fk_employeeId" db:"fk_employeeId"`
	FkComponentId	int 		`json:"fk_componentId" db:"fk_componentId"`
	FkCostCenterId	int 		`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type EditPayrollItems struct {
	Amount			*float64	`json:"amount" db:"amount"`
	FkPayrollRunId	*int 		`json:"fk_payrollRunId" db:"fk_payrollRunId"`
	FkEmployeeId	*int 		`json:"fk_employeeId" db:"fk_employeeId"`
	FkComponentId	*int 		`json:"fk_componentId" db:"fk_componentId"`
	FkCostCenterId	*int 		`json:"fk_costCenterId" db:"fk_costCenterId"`
}