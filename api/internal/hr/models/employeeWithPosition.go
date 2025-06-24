package models

type EmployeeWithPosition struct {
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"-"`
	Position string `db:"position_name" json:"position"`
}