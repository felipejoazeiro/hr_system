package hr

import "time"

type Employee struct {
	Matricula 		string 		`json: "matricula" binding:"required"`
	Nome 			string 		`json: "nome" binding:"required"`
	Email 			string		`json:"email" binding: "required"`
	Cargo			string		`json:"cargo" binding: "required"`
	Telefone 		string		`json:"telefone,omitempty"`
	DataNascimento	string		`json:"data_nascimento" binding:"required"`
	DataCadastro 	time.Time		`json:"data_cadastro"`

}