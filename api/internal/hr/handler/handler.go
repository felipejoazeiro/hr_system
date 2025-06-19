package handler 

import (
	"app/internal/hr/models"
	"app/internal/hr/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateEmployeeHandler(c *gin.Context){
	var newEmployee models.CreateEmployeeModel

	if err := c.ShouldBindJSON(&newEmployee); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savedEmployee, err := repository.CreateEmployee(newEmployee)

	if err != nil{
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, savedEmployee)
}