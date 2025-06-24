package repository

import (
	//"errors"
	"fmt"
	"strings"

	//"github.com/jmoiron/sqlx"
	"app/internal/hr/models"
	"app/pkg/db"

	"github.com/jmoiron/sqlx"

	"golang.org/x/crypto/bcrypt"
)

func CreateEmployee(e models.CreateEmployeeModel) (models.EmployeeModel, error) {
	tx, err := db.DB.Beginx()
	if err != nil {
		return models.EmployeeModel{}, err
	}
	query := `
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

	var result models.EmployeeModel
	stmt, err := db.DB.PrepareNamed(query)
	if err != nil {
		tx.Rollback()
		return result, err
	}
	err = stmt.Get(&result, e)

	login, err := generateUniqueLogin(e.Name, tx)
	if err != nil {
		tx.Rollback()
		return result, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(e.Registration), bcrypt.DefaultCost)

	accessQuery := `INSERT INTO access_employee (employee_id, login, password) VALUES ($1,$2,$3)`

	if _, err := tx.Exec(accessQuery, result.ID, login, string(hashedPassword)); err != nil {
		tx.Rollback()
		return result, err
	}

	if err := tx.Commit(); err != nil {
		return result, err
	}

	return result, nil
}

func generateUniqueLogin(name string, tx *sqlx.Tx) (string, error) {
	parts := strings.Fields(strings.ToLower(name))
	if len(parts) < 2 {
		return "", fmt.Errorf("precisa ter pelo menos dois nomes")
	}

	baseLogin := fmt.Sprintf("%s_%s", parts[0], parts[1])
	login := baseLogin

	var count int
	query := `SELECT COUNT (*) FROM access_employee WHERE login = $1`

	if err := tx.Get(&count, query, login); err != nil {
		return "", err
	}

	i := 2

	for count > 0 && i < len(parts) {
		login = fmt.Sprintf("%s_%s", parts[0], parts[i])
		if err := tx.Get(&count, query, login); err != nil {
			return "", err
		}
		i++
	}
	if count > 0 {
		i := 1
		for {
			newLogin := fmt.Sprintf("%s_%d", baseLogin, i)
			if err := tx.Get(&count, query, newLogin); err != nil {
				return "", err
			}
			if count == 0 {
				login = newLogin
				break
			}
			i++
		}
	}
	return login, nil
}
