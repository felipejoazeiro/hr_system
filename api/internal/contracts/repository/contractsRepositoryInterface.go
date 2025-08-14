package repository

import (
	"app/internal/contracts/models"
	"database/sql"
)

type ContractsRepositoryInterface interface {
	GetContract(id int) (models.GetAllContracts, error)
	GetContractByID(id int) (models.Contract, error)
	GetAllContracts() ([]models.GetAllContracts, error)
	CreateContract(tx *sql.Tx, input models.CreateContract) (int, error)
	EditContract(id int, input models.EditContract) (int, error)
	DeactivateContract(id int) error
	ReactivateContract(id int) error

	GetContractInfos(id int) (models.ContractInfoModel, error)
	CreateContractInfos(tx *sql.Tx, input models.CreateContractInfo) (int, error)
	EditContractInfos(id int, input models.EditContractInfo) (int, error)

	GetContractDates(id int) (models.ContractDatesModel, error)
	CreateContractDates(tx *sql.Tx, input models.CreateContractDates) (int, error)
	EditContractDates(id int, input models.EditContractDates) (int, error)

	GetContractDiscount(id int) (models.ContractDiscountModel, error)
	CreateContractDiscount(tx *sql.Tx, input models.CreateContractDiscount) (int, error)
	EditContractDiscount(id int, input models.EditContractDiscount) (int, error)

	GetContractRhInfo(id int) (models.ContractRhInfoModel, error)
	CreateContractRhInfo(tx *sql.Tx, input models.CreateContractRhInfo) (int, error)
	EditContractRhInfo(id int, input models.EditContractRhInfo) (int, error)

	GetContractValues(id int) (models.ContractValuesModel, error)
	CreateContractValues(tx *sql.Tx, input models.CreateContractValues) (int, error)
	EditContractValues(id int, input models.EditContractValues) (int, error)
}
