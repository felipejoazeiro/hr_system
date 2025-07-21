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

func (h *SindicateHandler) EditSindicateValues(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	idToUpdate, err := h.repo.getSindicateValuesId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Sindicate Values ID"})
		return
	}
	
	var input models.EditSindicateValues
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}
	sindicate, err := h.repo.EditSindicateValues(idToUpdate, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicate)
}

func (h *SindicateHandler) EditSindicatePercents(c *gin.Context) { 
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	idToUpdate, err := h.repo.getSindicatePercentsId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Sindicate Percents ID"})
		return
	}

	var input models.EditSindicatePercents
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}

	sindicatePercents, err := h.repo.EditSindicatePercents(idToUpdate, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicatePercents)
}

func (h *SindicateHandler) EditSindicateAuthorization(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	idToUpdate, err := h.repo.getSindicateAuthorizationId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Sindicate Authorization ID"})
		return
	}

	var input models.EditSindicateAuthorization
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}

	sindicateAuthorization, err := h.repo.EditSindicateAuthorization(idToUpdate, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicateAuthorization)
}

func (h *SindicateHandler) EditSindicateBreakfest(c *gin.Context) { 
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	idToUpdate, err := h.repo.getSindicateBreakfestId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Sindicate Breakfest ID"})
		return
	}

	var input models.EditSindicateBreakfest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}

	sindicateBreakfest, err := h.repo.EditSindicateBreakfest(idToUpdate, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicateBreakfest)
}

func (h *SindicateHandler) EditSindicateCarePackage(c *gin.Context) { 
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	idToUpdate, err := h.repo.getSindicateCarePackageId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Sindicate CarePackage ID"})
		return
	}

	var input models.EditSindicateCarePackage
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}

	sindicateCarePackage, err := h.repo.EditSindicateCarePackage(idToUpdate, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicateCarePackage)
}

func (h *SindicateHandler) EditSindicateVaucher(c *gin.Context) { 
	paramId := c.Param("id")
	id, err := strcov.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.h{"error": "Invalid ID"})
		return
	}

	idToUpdate, err := h.repo.getSindicateVoucherId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Sindicate Voucher ID"})
		return
	}

	var input models.EditSindicateVoucher
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Format"})
		return
	}

	sindicateVoucher, err := h.repo.EditSindicateVoucher(idToUpdate, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sindicateVoucher)
}

func (h *SindicateHandler) DeactivateSindicate(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivateSindicate(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sindicate deactivated sucessfully"})
}

func (h *SindicateHandler) ReactivateSindicate(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateSindicate(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sindicate reactivated successfully"})
}