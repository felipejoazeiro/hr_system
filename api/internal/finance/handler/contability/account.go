package handler

type AccountsHandler struct {
	repo repository.FinanceContabilityRepository
}

func NewAccountsHandler(repo repository.FinanceContabilityRepository) *AccountsHandler {
	return &AccountsHandler{
		repo: repo,
	}
}

func (h *AccountsHandler) GetAllAccounts(c *gin.Context) {
	accounts, err := h.repo.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (h *AccountsHandler) CreateAccount(c *gin.Context) {
	var account models.CreateAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *AccountsHandler) GetAccountByID(c *gin.Context) {
	id := c.Param("id")
	account, err := h.repo.GetAccountByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *AccountsHandler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var account models.EditAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateAccount(id, account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully"})
}

func (h *AccountsHandler) DeactivateAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deactivated successfully"})
}

func (h *AccountsHandler) ReactivateAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateAccount(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account reactivated successfully"})
}