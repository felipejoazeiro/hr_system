package handler

type PayrollComponentsHandler struct {
	repo repository.FinancePayrollRepository
}

func NewPayrollComponentsHandler(repo repository.FinancePayrollRepository) *PayrollComponentsHandler {
	return &PayrollComponentsHandler{
		repo: repo,
	}
}

func (h *PayrollComponentsHandler) GetAllPayrollComponents(c *gin.Context) {
	payrollComponents, err := h.repo.GetAllPayrollComponents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payrollComponents)
}

func (h *PayrollComponentsHandler) GetPayrollComponentsByID(c *gin.Context) {
	id := c.Param("id")
	payrollComponent, err := h.repo.GetPayrollComponentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payrollComponent)
}

func (h *PayrollComponentsHandler) CreatePayrollComponent(c *gin.Context) {
	var payrollComponent models.CreatePayrollComponent
	if err := c.ShouldBindJSON(&payrollComponent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreatePayrollComponent(payrollComponent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *PayrollComponentsHandler) UpdatePayrollComponent(c *gin.Context) {
	id := c.Param("id")
	var payrollComponent models.EditPayrollComponent
	if err := c.ShouldBindJSON(&payrollComponent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdatePayrollComponent(id, payrollComponent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payroll component updated successfully"})
}

func (h *PayrollComponentsHandler) DeactivatePayrollComponent(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.DeactivatePayrollComponent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payroll component deactivated successfully"})
}

func (h *PayrollComponentsHandler) ReactivatePayrollComponent(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.ReactivatePayrollComponent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payroll component reactivated successfully"})
}