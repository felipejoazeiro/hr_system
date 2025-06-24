package handler

import (
	"app/internal/hr/repository"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployeeByIDHandler(c *gin.Context){
	id := c.Param("id")

	employee, err := repository.Repository.GetEmployeeByID(id)
	if err != nil{
		handleEmployeeError(c,err)
	}

	c.JSON(http.StatusOK, employee)
}

func GetEmployeeByRegistration(c *gin.Context){
	registration := c.Param("registration")

	employee, err := repository.Repository.GetEmployeeByRegistration(registration)
	if err != nil{
		handleEmployeeError(c, err)
	}
	c.JSON(http.StatusOK, employee)
}

func handleEmployeeError(c *gin.Context, err error){
	if errors.Is(err, sql.ErrNoRows){
		c.JSON(http.StatusNotFound, gin.H{"error": "Funcionário não encontrado"})
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Erro ao buscar funcionário",
			"details": err.Error(),
		})
	}
}