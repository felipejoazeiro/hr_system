package handler

type FinanceApInvoiceItemHandler struct {
	repo repository.FinanceApInvoiceItemRepository
}

func NewFinanceApInvoiceItemHandler(repo repository.FinanceApInvoiceItemRepository) *FinanceApInvoiceItemHandler {
	return &FinanceApInvoiceItemHandler{
		repo: repo,
	}
}

func (h *FinanceApInvoiceItemHandler) GetAllApInvoiceItems(c *gin.Context) {
	items, err := h.repo.GetAllApInvoiceItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *FinanceApInvoiceItemHandler) GetApInvoiceItems(c *gin.Context) {
	id := c.Param("id")
	item, err := h.repo.GetApInvoiceItems(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *FinanceApInvoiceItemHandler) CreateApInvoiceItems(c *gin.Context) {
	var item models.CreateApInvoiceItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.repo.CreateApInvoiceItems(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *FinanceApInvoiceItemHandler) EditApInvoiceItems(c *gin.Context) {
	id := c.Param("id")
	var item models.EditApInvoiceItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.EditApInvoiceItems(id, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApInvoiceItemHandler) DeactivateApInvoiceItems(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateApInvoiceItems(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApInvoiceItemHandler) ReactivateApInvoiceItems(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateApInvoiceItems(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
