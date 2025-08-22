package handler 


type FinanceArHandler struct {
	repo repository.FinanceArRepository
}

func NewFinanceArHandler(repo repository.FinanceArRepository) *FinanceArHandler {
	return &FinanceArHandler{
		repo: repo,
	}
}

// INITIAL CRUD CUSTOMERS 
func (h *FinanceArHandler) GetCustomersById(c *gin.Context) {
	customerId := c.Param("id")
	customer, err := h.repo.GetCustomerById(customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *FinanceArHandler) GetAllCustomers(c *gin.Context) {
	customers, err := h.repo.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *FinanceArHandler) CreateCustomer(c *gin.Context) {
	var customer models.CreateCustomer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return 
	}
	customerId, err := h.repo.CreateCustomers(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": customerId})
}

func (h *FinanceArHandler) UpdateCustomer(c *gin.Context){
	var customer models.UpdateCustomer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	customerId, err := h.repo.UpdateCustomers(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": customerId})
}

func (h *FinanceArHandler) DeactivateCustomer(c *gin.Context){
	customerId := c.Param("id")
	err := h.repo.DeactivateCustomer(customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *FinanceArHandler) ReactivateCustomer(c *gin.Context) {
	customerId := c.Param("id")
	err := h.repo.ReactivateCustomer(customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customerId)
}