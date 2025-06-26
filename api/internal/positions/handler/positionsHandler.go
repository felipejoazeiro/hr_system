package handler


type PositionsHandler struct {
	repo *repository.PositionRepository
}

func NewPositionHandler(repo *repository.PositionRepository) *PositionHandler{
	return &PositionHandler{repo: repo}
}

func (h *PositionHandler) CreatePosition(c *gin.Context){
	var input models.CreatePosition
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato Inválido"})
		return
	}
	position, err := h.repo.CreatePosition(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated,position)
}

func (h *PositionHandler) GetAllPositions(c *gin.Context){
	positions, err := h.repo.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, positions)
}

func (h *PositionHandler) EditPosition(c *gin.Context){}

func (h *PositionHandler) DeactivatePosition(c *gin.Context){}

func (h *PositionHandler) ReactivePosition(c *gin.Context){}