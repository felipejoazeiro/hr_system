package models

type GetAllContracts struct {
	ID				int		`json:"id" db:"id"`
	Name			string	`json:"name" db:"name"`
	Code			string	`json:"code" db:"code"`
	InitialDate		time	`json:"initial_date" db:"initial_date"`	
}

type GetContract struct {
	ID				int			`json:"id" db:"id"`
	Name			string		`json:"name" db:"name"`
	Code			string		`json:"code" db:"code"`
	Employee		string		`json:"employee" db:"employee"`
	Research		bool		`json"research" db:"research"`
	IsActive		bool		`json:"is_active" db:"is_active"`
	Cpf 			string		`json:"cpf" db:"cpf"`
	Address			string		`json:"address" db:"address"`
	NRO				string		`json:"nro" db:"nro"`
	State			string		`json:"state" db:"state"`
	City 			string 		`json:"city" db:"city"`	
	InitialDate		string		`json:"initial_date" db:"initial_date"`
	LimitDate		string		`json:"limit_date" db:"limit_date"`
	
}