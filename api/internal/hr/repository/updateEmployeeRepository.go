package repository

import (
	"app/internal/hr/models"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func (r *EmployeeRepository) UpdateEmployeeRepository(employeeID string, update models.EditEmployeeModel) (*models.CreateEmployeeModel, error) {
	setClause, args := buildUpdateQuery(update)

	if setClause == "" {
		return nil, fmt.Errorf("não há campos para atualizar")
	}
	query := "UPDATE employees SET " + setClause + " WHERE id = ? RETURNING *"
	args = append(args, employeeID)

	var employee models.CreateEmployeeModel
	err := r.db.QueryRow(query, args...).Scan(
		&employee.ActiveEmployee,
		&employee.Registration,
		&employee.Address,
		&employee.BirthDate,
		&employee.CEP,
		&employee.CPF,
		&employee.CardNumber,
		&employee.City,
		&employee.CodeLine,
		&employee.CodeOperator,
		&employee.ContractDate,
		&employee.District,
		&employee.ElectoralSection,
		&employee.ElectoralZone,
		&employee.Email,
		&employee.EntryDate,
		&employee.FkCargo,
		&employee.FkTicketType,
		&employee.Identification,
		&employee.Name,
		&employee.NroRG,
		&employee.NroRG,
		&employee.NroTitle,
		&employee.NroWorkCard,
		&employee.Phone,
		&employee.Photo,
		&employee.QttDailyTicket,
		&employee.Registration,
		&employee.SeriesWorkCard,
		&employee.StateRG,
		&employee.States,
		&employee.TicketValue,
		&employee.UF,
	)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func buildUpdateQuery(update models.EditEmployeeModel) (string, []interface{}) {
	var sets []string
	var args []interface{}

	// Campos básicos
	addField(&sets, &args, "active_employee", update.ActiveEmployee)
	addField(&sets, &args, "registration", update.Registration)
	addField(&sets, &args, "name_employee", update.Name)
	addField(&sets, &args, "email", update.Email)

	// Datas
	addField(&sets, &args, "entry_date", update.EntryDate)
	addField(&sets, &args, "contract_date", update.ContractDate)

	// Informações pessoais
	addField(&sets, &args, "photo", update.Photo)
	addField(&sets, &args, "birthdate", update.BirthDate)
	addField(&sets, &args, "fk_cargo", update.FkCargo)

	// Documentos
	addField(&sets, &args, "nro_title", update.NroTitle)
	addField(&sets, &args, "electoral_zone", update.ElectoralZone)
	addField(&sets, &args, "electoral_section", update.ElectoralSection)
	addField(&sets, &args, "nro_rg", update.NroRG)
	addField(&sets, &args, "state_rg", update.StateRG)
	addField(&sets, &args, "nro_work_card", update.NroWorkCard)
	addField(&sets, &args, "series_work_card", update.SeriesWorkCard)
	addField(&sets, &args, "cpf", update.CPF)

	// Contato e endereço
	addField(&sets, &args, "phone", update.Phone)
	addField(&sets, &args, "address", update.Address)
	addField(&sets, &args, "district", update.District)
	addField(&sets, &args, "city", update.City)
	addField(&sets, &args, "states", update.States)
	addField(&sets, &args, "uf", update.UF)
	addField(&sets, &args, "cep", update.CEP)

	// Identificação e tickets
	addField(&sets, &args, "identification", update.Identification)
	addField(&sets, &args, "code_operator", update.CodeOperator)
	addField(&sets, &args, "code_line", update.CodeLine)
	addField(&sets, &args, "card_number", update.CardNumber)
	addField(&sets, &args, "qtt_daily_ticker", update.QttDailyTicket)
	addField(&sets, &args, "ticket_value", update.TicketValue)
	addField(&sets, &args, "fk_ticket_type", update.FkTicketType)

	// Adiciona updated_at
	sets = append(sets, "updated_at = ?")
	args = append(args, time.Now())

	return strings.Join(sets, ", "), args
}

// Função genérica para adicionar campos ao SET
func addField(sets *[]string, args *[]interface{}, field string, value interface{}) {
	if !isNil(value) {
		*sets = append(*sets, field+" = ?")
		*args = append(*args, value)
	}
}

// Verifica se um valor é nil (funciona com interfaces e ponteiros)
func isNil(val interface{}) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
