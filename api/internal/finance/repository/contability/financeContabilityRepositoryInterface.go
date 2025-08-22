package repository

import "api/internal/finance/models"

type FinanceContabilityRepository interface {
	GetAllCostCenters() ([]models.CostCenters, error)
	CreateCostCenter(costCenter models.CreateCostCenters) (int, error)
	GetCostCenterByID(id int) (models.CostCenters, error)
	UpdateCostCenter(id int, costCenter models.EditCostCenters) error
	DeactivateCostCenter(id int) error
	ReactivateCostCenter(id int) error

	GetAllAccounts() ([]models.Account, error)
	CreateAccount(account models.CreateAccount) (int, error)
	GetAccountByID(id int) (models.Account, error)
	UpdateAccount(id int, account models.EditAccount) error
	DeactivateAccount(id int) error
	ReactivateAccount(id int) error

	GetAllBankAccounts() ([]models.BankAccount, error)
	CreateBankAccount(account models.CreateBankAccount) (int, error)
	GetBankAccountByID(id int) (models.BankAccount, error)
	UpdateBankAccount(id int, account models.EditBankAccount) error
	DeactivateBankAccount(id int) error
	ReactivateBankAccount(id int) error

	GetAllBankTransactions() ([]models.BankTransaction, error)
	CreateBankTransaction(transaction models.CreateBankTransaction) (int, error)
	GetBankTransactionByID(id int) (models.BankTransaction, error)
	UpdateBankTransaction(id int, transaction models.EditBankTransaction) error
	DeactivateBankTransaction(id int) error
	ReactivateBankTransaction(id int) error
}