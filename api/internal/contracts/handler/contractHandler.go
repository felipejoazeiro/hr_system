package handler

import (
	"app/internal/contracts/models"
	"app/internal/contracts/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContractHandler struct {
	repo repository.ContractsRepositoryInterface
}

func NewContractHandler(repo repository.ContractsRepositoryInterface) *ContractHandler {
	return &ContractHandler{repo: repo}
}

func (h *ContractHandler) CreateContract(c *gin.Context) {
	var input models.CreateFullContractRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
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

func (h *ContractHandler) EditContractValues(c *gin.Context) {
	id := c.Param("id")
	var input models.EditContractValues

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	updated, err := h.repo.EditContractValues(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) EditContractDates(c *gin.Context) {
	id := c.Param("id")

	var input models.EditContractDates

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	updated, err := h.repo.EditContractDates(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) EditContractDiscount(c *gin.Context) {
	id := c.Param("id")

	var input models.EditContractDiscount

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}
	updated, err := h.repo.EditContractDiscount(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) EditContractInfo(c *gin.Context) {
	id := c.Param("id")

	var input models.EditContractInfo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}
	updated, err := h.repo.EditContractInfo(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) EditContractRhInfo(c *gin.Context) {
	id := c.Param("id")

	var input models.EditContractRhInfo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}
	updated, err := h.repo.EditContractRhInfo(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) EditContractValues(c *gin.context) {
	id := c.Param("id")

	var input models.EditContractValues

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}
	updated, err := h.repo.EditContractValues(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) DeactivateContract(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.DeactivateContract(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contract deactivated successfully"})
}

func (h *ContractHandler) ReactivateContract(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.ReactivateContract(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contract reactivated successfully"})
}
