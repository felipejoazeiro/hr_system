package models

type DepartmentModel struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Code      string `json:"code" db:"code"`
	ManagerId string `json:"manager_id" db:"manager_id"`
	IsActive  bool   `json:"is_active" db:"is_active"`
}

type GetDepartmentsRequest struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Code        string `json:"code" db:"code"`
	ManagerName string `json:"manager_name" db:"manager_name"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

type CreateDepartmentRequest struct {
	Name      string `json:"name" db:"name"`
	Code      string `json:"code" db:"code"`
	ManagerId *int   `json:"manager_id" db:"manager_id"`
}

type EditDepartmentRequest struct {
	Name      *string `json:"name" db:"name"`
	Code      *string `json:"code" db:"code"`
	ManagerId *string `json:"manager_id" db:"manager_id"`
	IsActive  *bool   `json:"is_active" db:"is_active"`
}

type RemoveDepartmentRequest struct {
	ID   string `json:"id" db:"id"`
	Code string `json:"code" db:"code"`
}
