package handler 


type FinanceArHandler struct {
	repo repository.FinanceArRepository
}

func NewFinanceArHandler(repo repository.FinanceArRepository) *FinanceArHandler {
	return &FinanceArHandler{
		repo: repo,
	}
}

// INITIAL CRUD ACCOUNTS RECEIVABLE ITEMS
func (h *FinanceArHandler) GetAllArInvoicesItems(c *gin.Context) {
	invoicesItems, err := h.repo.GetAllArInvoicesItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoicesItems)
}

func (h *FinanceArHandler) GetArInvoiceItemById(c *gin.Context) {
	itemId := c.Param("id")
	item, err := h.repo.GetArInvoiceItemById(itemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *FinanceArHandler) CreateArInvoiceItem(c *gin.Context) {
	var item models.CreateArInvoiceItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	itemId, err := h.repo.CreateArInvoiceItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": itemId})
}

func (h *FinanceArHandler) UpdateArInvoiceItem(c *gin.Context) {
	var item models.UpdateArInvoiceItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	itemId, err := h.repo.UpdateArInvoiceItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": itemId})
}

func (h *FinanceArHandler) DeactivateArInvoiceItem(c *gin.Context) {
	itemId := c.Param("id")
	if err := h.repo.DeactivateArInvoiceItem(itemId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item desativado com sucesso"})
}

func (h *FinanceArHandler) ReactivateArInvoiceItem(c *gin.Context) {
	itemId := c.Param("id")
	if err := h.repo.ReactivateArInvoiceItem(itemId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item reativado com sucesso"})
}