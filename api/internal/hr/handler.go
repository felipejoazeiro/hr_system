package hr 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateEmployeeHandler(c *gin.Context){
	var newEmployee Employee

	if err := c.ShouldBinJSON(&newEmployee); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savedEmployee, err := CreateEmployee(newEmployee)

	if err != nil{
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, savedEmployee)
}