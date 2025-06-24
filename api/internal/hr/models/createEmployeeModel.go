package models

import "time"

type CreateEmployeeModel struct {
	ActiveEmployee   bool      `db:"active_employee" json:"active_employee"`
	Registration     string    `db:"registration" json:"registration"`
	Name             string    `db:"name_employee" json:"name_employee"`
	Email            string    `db:"email" json:"email"`
	EntryDate        time.Time `db:"entry_date" json:"entry_date"`
	ContractDate     time.Time `db:"contract_date" json:"contract_date"`
	Photo            string    `db:"photo" json:"photo"`
	BirthDate        string    `db:"birthdate" json:"date"` 
	FkCargo          int       `db:"fk_cargo" json:"fk_cargo"`
	NroTitle         string    `db:"nro_title" json:"nro_title"`
	ElectoralZone    string    `db:"electoral_zone" json:"electoral_zone"`
	ElectoralSection string    `db:"electoral_section" json:"electoral_section"`
	NroRG            int       `db:"nro_rg" json:"nro_rg"`
	StateRG          string    `db:"state_rg" json:"state_rg"`
	NroWorkCard      int       `db:"nro_work_card" json:"nro_work_card"`
	SeriesWorkCard   string    `db:"series_work_card" json:"series_work_card"`
	CPF              int       `db:"cpf" json:"cpf"`
	Phone            int       `db:"phone" json:"phone"`
	Address          string    `db:"address" json:"address"`
	District         string    `db:"district" json:"district"`
	City             string    `db:"city" json:"city"`
	States           string    `db:"states" json:"states"`
	UF               string    `db:"uf" json:"uf"`
	CEP              int       `db:"cep" json:"cep"`
	Identification   int       `db:"identification" json:"identification"`
	CodeOperator     string    `db:"code_operator" json:"code_operator"`
	CodeLine         string    `db:"code_line" json:"code_line"`
	CardNumber       int       `db:"card_number" json:"card_number"`
	QttDailyTicket   int       `db:"qtt_daily_ticker" json:"qtt_daily_ticker"`
	TicketValue      int       `db:"ticket_value" json:"ticket_value"`
	FkTicketType     int       `db:"fk_ticket_type" json:"fk_ticket_type"`
}