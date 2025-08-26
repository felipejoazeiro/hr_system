package handler

type EmployeeSalaryHistoryHandler struct {
	repo repository.FinancePayrollRepository
}

func NewEmployeeSalaryHistoryHandler(repo repository.FinancePayrollRepository) *EmployeeSalaryHistoryHandler {
	return &EmployeeSalaryHistoryHandler{
		repo: repo,
	}
}

func (h *EmployeeSalaryHistoryHandler) GetAllEmployeeSalaryHistory(c *gin.Context) {
	employeeSalaryHistories, err := h.repo.GetAllEmployeeSalaryHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employeeSalaryHistories)
}

func (h *EmployeeSalaryHistoryHandler) GetEmployeeSalaryHistoryByID(c *gin.Context) {
	id := c.Param("id")
	employeeSalaryHistory, err := h.repo.GetEmployeeSalaryHistoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employeeSalaryHistory)
}

func (h *EmployeeSalaryHistoryHandler) CreateEmployeeSalaryHistory(c *gin.Context) {
	var employeeSalaryHistory models.CreateEmployeeSalaryHistory
	if err := c.ShouldBindJSON(&employeeSalaryHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateEmployeeSalaryHistory(employeeSalaryHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *EmployeeSalaryHistoryHandler) UpdateEmployeeSalaryHistory(c *gin.Context) {
	id := c.Param("id")
	var employeeSalaryHistory models.EditEmployeeSalaryHistory
	if err := c.ShouldBindJSON(&employeeSalaryHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateEmployeeSalaryHistory(id, employeeSalaryHistory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee salary history updated successfully"})
}
