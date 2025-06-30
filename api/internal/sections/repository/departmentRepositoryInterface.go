package repository

import (
	"app/internal/sections/models"
)

type DepartmentRepositoryInterface interface {
    GetAllDepartments() ([]models.GetDepartmentsRequest, error)
    CreateDepartment(input models.CreateDepartmentRequest) (models.DepartmentModel, error)
    EditDepartment(id string, input models.EditDepartmentRequest) (models.DepartmentModel, error) // Adicionado id como parâmetro
    DeactivateDepartment(id string) error
    ReactivateDepartment(id string) error
}
