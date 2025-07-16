package handler


type SindicateHandler struct {
	repo repository.SindicateRepositoryInterface
}

func NewSindicateHandler(repo  repository.SindicateRepositoryInterface) *SindicateHandler {
	return &SindicateHandler{repo: repo}
}

func (h *SindicateHandler) GetAllSindicates(c *gin.Context) {
	sindicates, err := h.repo.GetAllSindicates()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sindicates)
}

func (h *SindicateHandler) CreateSindicate(c *gin.Context) {
	var input models.CreateSindicateReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato Inválido"})
		return
	}
	sindicate, err := h.repo.CreateSindicate(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicate)
}

func (h *SindicateHandler) EditSindicate(c *gin.Context){
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	var input models.EditSindicate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	updatedId, err := h.repo.UpdateSindicate(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update sindicate error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "id": "updatedId"})
}