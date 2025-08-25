package models

type EmployeeDataForList struct {
	ID         int `json:"id" db:"id"`
	Register   string `json:"register" db:"register"`
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Position   string `json:"position" db:"position"`
	Department string `json:"department" db:"department"`
	Active     bool   `json:"active" db:"active"`
}

type EmployeeFilter struct {
	Position        *string `form:"position"`
	Department      *string `form:"department"`
	IncludeInactive *bool   `form:"include_inactive"`
}
