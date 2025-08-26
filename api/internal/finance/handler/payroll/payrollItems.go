package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your_project/models"
	"your_project/repository"
)

type PayrollItemsHandler struct {
	repo repository.FinancePayrollRepository
}

func NewPayrollItemsHandler(repo repository.FinancePayrollRepository) *PayrollItemsHandler {
	return &PayrollItemsHandler{
		repo: repo,
	}
}

func (h *PayrollItemsHandler) GetAllPayrollItems(c *gin.Context) {
	payrollItems, err := h.repo.GetAllPayrollItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payrollItems)
}

func (h *PayrollItemsHandler) GetPayrollItemByID(c *gin.Context) {
	id := c.Param("id")
	payrollItem, err := h.repo.GetPayrollItemByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payrollItem)
}

func (h *PayrollItemsHandler) CreatePayrollItem(c *gin.Context) {
	var payrollItem models.CreatePayrollItem
	if err := c.ShouldBindJSON(&payrollItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreatePayrollItem(payrollItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *PayrollItemsHandler) UpdatePayrollItem(c *gin.Context) {
	id := c.Param("id")
	var payrollItem models.EditPayrollItem
	if err := c.ShouldBindJSON(&payrollItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdatePayrollItem(id, payrollItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payroll item updated successfully"})
}
