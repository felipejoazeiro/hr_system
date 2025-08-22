package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BankAccountsHandler struct {
	repo repository.FinanceContabilityRepository
}

func NewBankAccountsHandler(repo repository.FinanceContabilityRepository) *BankAccountsHandler {
	return &BankAccountsHandler{
		repo: repo,
	}
}

func (h *BankAccountsHandler) GetAllBankAccounts(c *gin.Context) {
	bankAccounts, err := h.repo.GetAllBankAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bankAccounts)
}

func (h *BankAccountsHandler) CreateBankAccount(c *gin.Context) {
	var bankAccount models.CreateBankAccount
	if err := c.ShouldBindJSON(&bankAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateBankAccount(bankAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *BankAccountsHandler) GetBankAccountByID(c *gin.Context) {
	id := c.Param("id")
	bankAccount, err := h.repo.GetBankAccountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bankAccount)
}

func (h *BankAccountsHandler) UpdateBankAccount(c *gin.Context) {
	id := c.Param("id")
	var bankAccount models.EditBankAccount
	if err := c.ShouldBindJSON(&bankAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateBankAccount(id, bankAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank account updated successfully"})
}

func (h *BankAccountsHandler) DeactivateBankAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateBankAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank account deactivated successfully"})
}

func (h *BankAccountsHandler) ReactivateBankAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateBankAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank account reactivated successfully"})
}
