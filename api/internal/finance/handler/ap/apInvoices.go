package handler

type FinanceArHandler struct {
	repo repository.FinanceArRepository
}

func NewFinanceArHandler(repo repository.FinanceArRepository) *FinanceArHandler {
	return &FinanceArHandler{
		repo: repo,
	}
}

func (h *FinanceApHandler) GetAllApInvoices(c *gin.Context) {
	invoices, err := h.repo.GetAllApInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func (h *FinanceApHandler) GetApInvoices(c *gin.Context) {
	id := c.Param("id")
	invoice, err := h.repo.GetApInvoices(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (h *FinanceApHandler) CreateApInvoices(c *gin.Context) {
	var invoice models.CreateApInvoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.repo.CreateApInvoices(invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *FinanceApHandler) EditApInvoices(c *gin.Context) {
	id := c.Param("id")
	var invoice models.EditApInvoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.EditApInvoices(id, invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApHandler) DeactivateApInvoices(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateApInvoices(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApHandler) ReactivateApInvoices(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateApInvoices(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
