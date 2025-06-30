package models

type PositionModel struct {
	ID			string	`json:"id" db:"id"`
	Name		string		`json:"name" db:"name"`
	Code    	string 		`json:"code" db:"code"`
	IsActive 	bool 		`json:"is_active" db:"is_active"`
}

type CreatePosition struct {
	Name	string			`json:"name" db:"name"`
	Code 	string			`json:"code" db:"code"`
}

type EditPosition struct {
	Name	*string			`json:"name" db:"name"`
	Code	*string			`json:"code" db:"code"`
}