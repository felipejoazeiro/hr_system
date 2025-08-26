package handler

type CostCentersHandler struct {
	repo repository.FinanceContabilityRepository
}

func NewCostCentersHandler(repo repository.FinanceContabilityRepository) *CostCentersHandler {
	return &CostCentersHandler{
		repo: repo,
	}
}

func (h *CostCentersHandler) GetAllCostCenters(c *gin.Context){
	costCenters, err := h.repo.GetAllCostCenters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	c.JSON(http.StatusOK, costCenters)
}

func (h *CostCentersHandler) GetCostCenterByID(c *gin.Context) {
	id := c.Param("id")
	costCenter, err := h.repo.GetCostCenterByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, costCenter)
}

func (h *CostCentersHandler) CreateCostCenter(c *gin.Context) {
	var costCenter models.CreateCostCenters
	if err := c.ShouldBindJSON(&costCenter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateCostCenter(costCenter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *CostCentersHandler) GetCostCenterByID(c *gin.Context) {
	id := c.Param("id")
	costCenter, err := h.repo.GetCostCenterByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, costCenter)
}

func (h *CostCentersHandler) UpdateCostCenter(c *gin.Context) {
	id := c.Param("id")
	var costCenter models.EditCostCenters
	if err := c.ShouldBindJSON(&costCenter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateCostCenter(id, costCenter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cost center updated successfully"})
}

func (h *CostCentersHandler) DeactivateCostCenter(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateCostCenter(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cost center deleted successfully"})
}

func (h *CostCentersHandler) ReactivateCostCenter(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateCostCenter(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cost center reactivated successfully"})
}