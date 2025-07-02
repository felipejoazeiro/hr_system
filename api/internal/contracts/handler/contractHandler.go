package handler

type ContractHandler struct {
	repo repository.ContractsRepositoryInterface
}

type NewContractHandler(repo repository.ContractsRepositoryInterface) *ContractHandler {
	return &ContractHandler{repo: repo}
}

func (h *ContractHandler) createContract(c *gin.Context) {
	var input models.CreateFullContractRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" :"JSON inválido: " + err.Error()})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("Erro ao iniciar transação:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno"})
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()


	infoId, err := h.repository.CreateContractInfo(tx, input.Info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar info do contrato"})
		return
	}

	datesId, err := h.repository.CreateContractDates(tx, input.Dates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar datas do contrato"})
		return
	}

	valuesId, err := h.repository.CreateContractValues(tx, input.Values)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar valores do contrato"})
		return
	}

	discountId, err := h.repository.CreateContractDiscount(tx, input.Discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar disconto do contrato"})
		return
	}

	rhId, err := h.repository.CreateContractRhInfo(tx, input.RhInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ai criar informações do RH do contrato"})
		return
	}

	contract := models.CreateFullContractRequest{
		Name:               input.Name,
		Code:               input.Code,
		Research:           input.Research,
		UsesCpf:            input.UsesCpf,
		IsActive:           input.IsActive,
		FkContractInfo:     infoId,
		FkContractDates:    datesId,
		FkContractValues:   valuesId,
		FkContractDiscount: discountId,
		FkContractRh:       rhId,
	}

	contractId, err := h.repo.CreateContract(tx, contract)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar contrato"})
		return
	}

	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao confirmar transação"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Contrato criado com sucesso", "id": contractID})

}