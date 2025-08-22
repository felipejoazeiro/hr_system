package handler 


type FinanceArHandler struct {
	repo repository.FinanceArRepository
}

func NewFinanceArHandler(repo repository.FinanceArRepository) *FinanceArHandler {
	return &FinanceArHandler{
		repo: repo,
	}
}

func (h *FinanceArHandler) GetAllArInvoices(c *gin.Context) {
	invoices, err := h.repo.GetAllArInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func (h *FinanceArHandler) GetInvoiceById(c *gin.Context) {
	invoiceId := c.Param("id")
	invoice, err := h.repo.GetArInvoiceById(invoiceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (h *FinanceArHandler) CreateArInvoices(c *gin.Context) {
	var invoice models.CreateArInvoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	invoiceId, err := h.repo.CreateArInvoices(invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": invoiceId})
}

func (h *FinanceArHandler) UpdateArInvoices(c *gin.Context) {
	var invoice models.UpdateArInvoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	invoiceId, err := h.repo.UpdateArInvoices(invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": invoiceId})
}

func (h *FinanceArHandler) DeactivateArInvoices(c *gin.Context) {
	invoiceId := c.Param("id")
	err := h.repo.DeactivateArInvoices(invoiceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *FinanceArHandler) ReactivateArInvoices(c *gin.Context) {
	invoiceId := c.Param("id")
	err := h.repo.ReactivateArInvoices(invoiceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoiceId)
}
