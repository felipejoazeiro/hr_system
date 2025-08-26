package repository

type FinancePayrollRepositoryInterface interface {

	GetAllPayrollComponents() ([]models.PayrollComponents, error)
	GetPayrollComponentByID(id int) (*models.PayrollComponents, error)
	CreatePayrollComponent(input models.CreatePayrollComponents) (int, error)
	EditPayrollComponent(input models.EditPayrollComponents) (int, error)
	DeactivatePayrollComponent(id int) error
	ReactivatePayrollComponent(id int) error

	GetAllEmployeeSalaryHistories() ([]models.EmployeeSalaryHistory, error)
	GetEmployeeSalaryHistoriesById(id int) (*models.EmployeeSalaryHistory, error)
	CreateEmployeeSalaryHistory(input models.CreateEmployeeSalaryHistory) (int, error)
	EditEmployeeSalaryHistory(input models.EditEmployeeSalaryHistory) (int, error)
	DeactivateEmployeeSalaryHistory(id int) error
	ReactivateEmployeeSalaryHistory(id int) error

	GetAllPayrollRuns() ([]models.PayrollRuns, error)
	GetPayrollRunByID(id int) (*models.PayrollRuns, error)
	CreatePayrollRun(input models.CreatePayrollRuns) (int, error)
	EditPayrollRun(input models.EditPayrollRuns) (int, error)
	DeactivatePayrollRun(id int) error
	ReactivatePayrollRun(id int) error

	GetAllPayrollItems() ([]models.PayrollItems, error)
	GetPayrollItemByID(id int) (*models.PayrollItems, error)
	CreatePayrollItem(input models.CreatePayrollItems) (int, error)
	EditPayrollItem(input models.EditPayrollItems) (int, error)
	DeactivatePayrollItem(id int) error
	ReactivatePayrollItem(id int) error

}