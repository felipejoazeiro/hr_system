package models

import "time"

type EmployeeModel struct {
	ID			 int	   `json:"id"`
	Registration string    `json:"registration"`
	Name         string    `json:"name_employee"`
	Email        string    `json:"email"`
	Role         string    `json:"cargo,omitempty"` // pode vir de join com tabela de cargos
	CreatedAt    time.Time `json:"created_at"`
}

type DeactivateRequest struct {
	Reason string `json:"reason"`
}
