package repository

import (
	"app/internal/sections/models"
)

type DepartmentRepositoryInterface interface {
    GetAllDepartments() ([]models.GetDepartmentsRequest, error)
    CreateDepartment(input models.CreateDepartmentRequest) (models.DepartmentModel, error)
    EditDepartment(id int, input models.EditDepartmentRequest) (models.DepartmentModel, error) // Adicionado id como parâmetro
    DeactivateDepartment(id int) error
    ReactivateDepartment(id int) error
}
