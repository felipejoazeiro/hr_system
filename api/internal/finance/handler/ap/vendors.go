package handler

type FinanceApHandler struct {
	repo repository.FinanceApRepository
}

func NewFinanceApHandler(repo repository.FinanceApRepository) *FinanceApHandler {
	return &FinanceApHandler{
		repo: repo,
	}
}

func (h *FinanceApHandler) GetVendor(c *gin.Context) {
	id := c.Param("id")
	vendor, err := h.repo.GetVendors(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vendor)
}

func (h *FinanceApHandler) GetAllVendors(c *gin.Context) {
	vendors, err := h.repo.GetAllVendors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vendors)
}

func (h *FinanceApHandler) CreateVendor(c *gin.Context)  {
	var vendor models.CreateVendors
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.repo.CreateVendors(vendor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *FinanceApHandler) EditVendor(c *gin.Context) {
	id := c.Param("id")
	var vendor models.EditVendors
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.EditVendors(id, vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApHandler) DeactivateVendor(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateVendor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *FinanceApHandler) ReactivateVendor(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateVendor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
