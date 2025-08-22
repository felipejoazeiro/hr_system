package handler 

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BankTransactionsHandler struct {
	repo repository.FinanceContabilityRepository
}

func NewBankTransactionsHandler(repo repository.FinanceContabilityRepository) *BankTransactionsHandler {
	return &BankTransactionsHandler{
		repo: repo,
	}
}

func (h *BankTransactionsHandler) GetAllBankTransactions(c *gin.Context) {
	bankTransactions, err := h.repo.GetAllBankTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bankTransactions)
}

func (h *BankTransactionsHandler) CreateBankTransaction(c *gin.Context) {
	var bankTransaction models.CreateBankTransaction
	if err := c.ShouldBindJSON(&bankTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateBankTransaction(bankTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *BankTransactionsHandler) GetBankTransactionByID(c *gin.Context) {
	id := c.Param("id")
	bankTransaction, err := h.repo.GetBankTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bankTransaction)
}

func (h *BankTransactionsHandler) UpdateBankTransaction(c *gin.Context) {
	id := c.Param("id")
	var bankTransaction models.EditBankTransaction
	if err := c.ShouldBindJSON(&bankTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateBankTransaction(id, bankTransaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank transaction updated successfully"})
}

func (h *BankTransactionsHandler) DeactivateBankTransaction(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateBankTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank transaction deactivated successfully"})
}

func (h *BankTransactionsHandler) ReactivateBankTransaction(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateBankTransaction(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bank transaction reactivated successfully"})
}
