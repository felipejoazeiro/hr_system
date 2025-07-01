package repository

import (
	"app/internal/contracts/models"
)

type ContractsRpositoryInterface interface {
	GetContract(id int) (models.GetContract, error)
	GetAllContracts() ([]models.GetAllContracts, error)
	CreateContract(input models.CreateContract) (models.ContractModel, error)
	EditContract(id int, input models.EditContract) (models.ContractoModel, error)
	DeactivateContract(id int) error
	ReactivateContract(id int) error

	GetContractDates(id int) (models.ContractDatesModel, error)
	CreateContractDates(input models.CreateContractDates) (models.ContractDatesModel, error)
	EditContractDate(id int, input models.EditContractDate) (models.ContractDatesModel, error)
	
	GetContractDiscount(id int) (models.ContractDiscount, error)
	CreateContractDiscount(input models.CreateContractDiscount) (models.ContractDiscountModel, error)
	EditContractDiscount(id int, input models.EditContractDiscount) (models.ContractDiscountModel, error)

	GetContractRhInfo(id int) (models.ContractRhInfoModels, error)
	CreateContractRhInfo(input models.CreateContractRhInfo) (models.ContractRhInfoModels, error)
	EditContractRhInfo(id int, input models.EditContractRhInfo) (models.ContractRhInfoModels, error)

	GetContractContact(id int) (models.ContractContact, error)
	CreateContractContact(input models.CreateContractContact) (models.ContractContactModel, error)
	EditContractContact(id int, input models.EditContractContact) (models.ContractContactModel, error)

	GetContractInfos(id int) (models.ContractInfoModel, error)
	CreateContractInfos(input models.CreateContractInfo) (models.ContractInfoModel, error) 
	EditContractContact(id int, input models.EditContractInfo) (models.ContractInfoModel, error)

	GetContractValues(id int) (models.ContractValuesModel, error)
	CreateContractValues(input models.CreateContractValues) (models.ContractValuesModel, error)
	EditContractValues(id int, models.EditContractValues) (models.ContractValuesModel, error)
}