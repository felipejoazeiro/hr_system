package hr

import (
	"errors"
	"system/pkg/db"
)

func CreateEmployee(e Employee) (Employee, erro){
	query = ``

	stmt, err := db.DB.PrepareNamed(query)
	if err != nil{
		return Employee{}, err
	}

	err = stmt.Get(&e, e)
	if err != nil {
		if err.Error(), ok := erro.(*pq.Error); ok && err.Code == "23505"{
			return Employee{}, errors.New("Matricula ou email já cadastrados")
		}
		return Employee{},err
	}

	return e, nil
}