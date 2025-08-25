package handler

import (
	"app/internal/contracts/models"
	"app/internal/contracts/repository"
	"app/pkg/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContractHandler struct {
	repo repository.ContractsRepositoryInterface
}

func NewContractHandler(repo repository.ContractsRepositoryInterface) *ContractHandler {
	return &ContractHandler{repo: repo}
}

func (h *ContractHandler) GetAllContracts(c *gin.Context) {
	contracts, err := h.repo.GetAllContracts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contratos"})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

func (h *ContractHandler) GetContractByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	contract, err := h.repo.GetContractByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contrato"})
		return
	}
	if contract.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado"})
		return
	}
	c.JSON(http.StatusOK, contract)
}

func (h *ContractHandler) EditContract(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var input models.EditContract
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}
	updated, err := h.repo.EditContract(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao editar contrato"})
		return
	}
	if updated == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contrato não encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contrato atualizado com sucesso", "id": updated})
}

func (h *ContractHandler) CreateContract(c *gin.Context) {
	var input models.CreateFullContractRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("Erro ao iniciar transação:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno"})
		return
	}

	// rollback por segurança (se Commit acontecer sem erro, o rollback não tem efeito)
	defer tx.Rollback()

	// protege contra panic
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	// Inserts auxiliares
	infoID, err := h.repo.CreateContractInfos(tx, input.Info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar info do contrato"})
		return
	}

	datesID, err := h.repo.CreateContractDates(tx, input.Dates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar datas do contrato"})
		return
	}

	valuesID, err := h.repo.CreateContractValues(tx, input.Values)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar valores do contrato"})
		return
	}

	discountID, err := h.repo.CreateContractDiscount(tx, input.Discount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar desconto do contrato"})
		return
	}

	rhID, err := h.repo.CreateContractRhInfo(tx, input.RhInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar informações de RH do contrato"})
		return
	}

	// Contrato principal
	contractID, err := h.repo.CreateContract(tx, models.CreateContract{
		Name:               input.Name,
		Code:               input.Code,
		Research:           input.Research,
		UsesCpf:            input.UsesCpf,
		IsActive:           input.IsActive,
		FkContractInfo:     infoID,
		FkContractDates:    datesID,
		FkContractValues:   valuesID,
		FkContractDiscount: discountID,
		FkContractRh:       rhID,
	})
	if err != nil {
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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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

func (h *ContractHandler) EditContractInfos(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input models.EditContractInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	updated, err := h.repo.EditContractInfos(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ContractHandler) EditContractRhInfo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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

func (h *ContractHandler) DeactivateContract(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.repo.DeactivateContract(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contrato desativado com sucesso"})
}

func (h *ContractHandler) ReactivateContract(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.repo.ReactivateContract(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contrato reativado com sucesso"})
}
