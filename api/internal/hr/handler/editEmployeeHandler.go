package handler

import (
	"app/internal/hr/models"
	"app/internal/hr/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditEmployeeHandler(c *gin.Context){
	employeeID:= c.Param("id")

	var updateData models.EditEmployeeModel
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}
	updatedEmployee, err := repository.Repository.UpdateEmployeeRepository(employeeID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar funcionário", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Funcionário atualizado com sucesso",
		"data": updatedEmployee,
	})
}