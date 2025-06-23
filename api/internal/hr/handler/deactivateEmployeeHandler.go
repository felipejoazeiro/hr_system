package handler

import (
	"app/internal/hr/models"
	"app/internal/hr/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeactivateEmployeeHandler(c *gin.Context){
	employeeID := c.Param("id")

	var request models.DeactivateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	err := repository.Repository.DeactivateEmployeeRepository(employeeID, request.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao desativar funcionário"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Funcionário desativado com sucesso",
		"employee_id": employeeID,
		"deactivation_reason": request.Reason,
	})
}