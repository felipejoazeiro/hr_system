package handler 

type FinanceApPaymentsHandler struct {
	repo repository.FinanceApPaymentsRepository
}

func NewFinanceApPaymentsHandler(repo repository.FinanceApPaymentsRepository) *FinanceApPaymentsHandler {
	return &FinanceApPaymentsHandler{
		repo: repo,
	}
}

func (h *FinanceApPaymentsHandler) GetAllApPayments(c *gin.Context) {
	payments, err := h.repo.GetAllApPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *FinanceApPaymentsHandler) GetApPayments(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.repo.GetApPayments(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *FinanceApPaymentsHandler) CreateApPayments(c *gin.Context) {
	var payment models.CreateApPayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.repo.CreateApPayments(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *FinanceApPaymentsHandler) EditApPayments(c *gin.Context) {
	id := c.Param("id")
	var payment models.EditApPayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.EditApPayments(id, payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApPaymentsHandler) DeactivateApPayments(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateApPayments(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApPaymentsHandler) ReactivateApPayments(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateApPayments(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}