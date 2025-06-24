package handler

import (
	"app/internal/hr/models"
	"app/internal/hr/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListEmployeeHandler(c *gin.Context) {
	var filter models.EmployeeFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetros inválidos"})
		return
	}

	employess, err := repository.Repository.ListEmployees(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar funcionários"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  employess,
		"count": len(employess),
	})
}
