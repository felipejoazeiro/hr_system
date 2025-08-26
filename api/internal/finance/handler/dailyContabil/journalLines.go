package handler

type JournalLinesHandler struct {
	repo repository.FinanceDailyContabilRepository
}

func NewJournalLinesHandler(repo repository.FinanceDailyContabilRepository) *JournalLinesHandler {
	return &JournalLinesHandler{
		repo: repo,
	}
}

func (h *JournalLinesHandler) GetAllJournalLines(c *gin.Context) {
	journalLines, err := h.repo.GetAllJournalLines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, journalLines)
}

func (h *JournalLinesHandler) GetJournalLineByID(c *gin.Context) {
	id := c.Param("id")
	journalLine, err := h.repo.GetJournalLineByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, journalLine)
}

func (h *JournalLinesHandler) CreateJournalLine(c *gin.Context) {
	var journalLine models.CreateJournalLine
	if err := c.ShouldBindJSON(&journalLine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateJournalLine(journalLine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *JournalLinesHandler) UpdateJournalLine(c *gin.Context) {
	id := c.Param("id")
	var journalLine models.EditJournalLine
	if err := c.ShouldBindJSON(&journalLine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.UpdateJournalLine(id, journalLine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Journal line updated successfully"})
}
