package handler

type JournalEntriesHandler struct {
	repo repository.FinanceDailyContabilRepository
}

func NewJournalEntriesHandler(repo repository.FinanceDailyContabilRepository) *JournalEntriesHandler {
	return &JournalEntriesHandler{
		repo: repo,
	}
}

func (h *JournalEntriesHandler) GetAllJournalEntries(c *gin.Context) {
	journalEntries, err := h.repo.GetAllJournalEntries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, journalEntries)
}

func (h *JournalEntriesHandler) GetJournalEntryByID(c *gin.Context) {
	id := c.Param("id")
	journalEntry, err := h.repo.GetJournalEntryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, journalEntry)
}

func (h *JournalEntriesHandler) CreateJournalEntry(c *gin.Context) {
	var entry models.CreateJournalEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateJournalEntry(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *JournalEntriesHandler) UpdateJournalEntry(c *gin.Context) {
	id := c.Param("id")
	var entry models.EditJournalEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateJournalEntry(id, entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Journal entry updated successfully"})
}
