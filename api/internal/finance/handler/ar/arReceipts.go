package handler 


type FinanceArHandler struct {
	repo repository.FinanceArRepository
}

func NewFinanceArHandler(repo repository.FinanceArRepository) *FinanceArHandler {
	return &FinanceArHandler{
		repo: repo,
	}
}

func (h *FinanceArHandler) GetAllArReceipts(c *gin.Context) {
	receipts, err := h.repo.GetAllArReceipts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receipts)
}

func (h *FinanceArHandler) GetArReceiptById(c *gin.Context) {
	receiptId := c.Param("id")
	receipt, err := h.repo.GetArReceiptById(receiptId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receipt)
}

func (h *FinanceArHandler) CreateArReceipt(c *gin.Context) {
	var receipt models.CreateArReceipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	receiptId, err := h.repo.CreateArReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": receiptId})
}

func (h *FinanceArHandler) UpdateArReceipt(c *gin.Context) {
	var receipt models.UpdateArReceipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	receiptId, err := h.repo.UpdateArReceipt(receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": receiptId})
}

func (h *FinanceArHandler) DeactivateArReceipt(c *gin.Context) {
	receiptId := c.Param("id")
	if err := h.repo.DeactivateArReceipt(receiptId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Receipt deactivated successfully"})
}

func (h *FinanceArHandler) ReactivateArReceipt(c *gin.Context) {
	receiptId := c.Param("id")
	if err := h.repo.ReactivateArReceipt(receiptId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Receipt reactivated successfully"})
}