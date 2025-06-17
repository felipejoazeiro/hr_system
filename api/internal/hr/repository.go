package hr

import (
	"errors"
	"system/pkg/db"
	"github.com/jmoiron/sqlx"
)

func CreateEmployee(e CreateEmployeeModel) (EmployeeModel, error){
	query = query := `
		INSERT INTO employees (
			active_employee, registration, name_employee, email,
			entry_date, contract_date, photo, date,
			fk_cargo, nro_title, electoral_zone, electoral_section,
			nro_rg, state_rg, nro_work_card, series_work_card,
			cpf, phone, address, district,
			city, states, uf, cep,
			identification, code_operator, code_line,
			card_number, qtt_daily_ticker, ticket_value, fk_ticket_type
		) VALUES (
			:active_employee, :registration, :name_employee, :email,
			:entry_date, :contract_date, :photo, :date,
			:fk_cargo, :nro_title, :electoral_zone, :electoral_section,
			:nro_rg, :state_rg, :nro_work_card, :series_work_card,
			:cpf, :phone, :address, :district,
			:city, :states, :uf, :cep,
			:identification, :code_operator, :code_line,
			:card_number, :qtt_daily_ticker, :ticket_value, :fk_ticket_type
		) RETURNING *
	`

	var result EmployeeModel

	stmt, err := db.DB.PrepareNamed(query)
	if err != nil {
		return result, err
	}
	err = stmt.Get(&result, e)

	return result, err
}