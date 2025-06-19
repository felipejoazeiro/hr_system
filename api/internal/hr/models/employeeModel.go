package models

import "time"

type EmployeeModel struct {
	ID           int       `json:"id"`
	Registration string    `json:"registration"`
	Name         string    `json:"name_employee"`
	Email        string    `json:"email"`
	Cargo        string    `json:"cargo,omitempty"` // pode vir de join com tabela de cargos
	EntryDate    time.Time `json:"entry_date"`
}