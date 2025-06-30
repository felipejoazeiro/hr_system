package handler

import (
	"app/internal/positions/models"
	"app/internal/positions/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	repo repository.PositionRepositoryInterface
}

func NewPositionHandler(repo repository.PositionRepositoryInterface) *PositionHandler {
	return &PositionHandler{repo: repo}
}

func (h *PositionHandler) CreatePosition(c *gin.Context) {
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
	c.JSON(http.StatusCreated, position)
}

func (h *PositionHandler) GetAllPositions(c *gin.Context) {
	positions, err := h.repo.GetAllPositions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, positions)
}

func (h *PositionHandler) EditPosition(c *gin.Context) {
	id := c.Param("id")

	var input models.EditPosition
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	pos, err := h.repo.EditPosition(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pos)
}

func (h *PositionHandler) DeactivatePosition(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeactivatePosition(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cargo desativado com sucesso"})

}

func (h *PositionHandler) ReactivePosition(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.ReactivePosition(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cargo reativado com sucesso!"})
}
