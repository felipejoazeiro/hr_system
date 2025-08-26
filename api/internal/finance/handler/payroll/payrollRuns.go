package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your_project/models"
	"your_project/repository"
)

type PayrollRunsHandler struct {
	repo repository.FinancePayrollRepository
}

func NewPayrollRunsHandler(repo repository.FinancePayrollRepository) *PayrollRunsHandler {
	return &PayrollRunsHandler{
		repo: repo,
	}
}

func (h *PayrollRunsHandler) GetAllPayrollRuns(c *gin.Context) {
	payrollRuns, err := h.repo.GetAllPayrollRuns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payrollRuns)
}

func (h *PayrollRunsHandler) GetPayrollRunByID(c *gin.Context) {
	id := c.Param("id")
	payrollRun, err := h.repo.GetPayrollRunByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payrollRun)
}

func (h *PayrollRunsHandler) CreatePayrollRun(c *gin.Context) {
	var payrollRun models.CreatePayrollRun
	if err := c.ShouldBindJSON(&payrollRun); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreatePayrollRun(payrollRun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *PayrollRunsHandler) UpdatePayrollRun(c *gin.Context) {
	id := c.Param("id")
	var payrollRun models.EditPayrollRun
	if err := c.ShouldBindJSON(&payrollRun); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdatePayrollRun(id, payrollRun); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payroll run updated successfully"})
}
