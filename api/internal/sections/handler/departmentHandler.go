package handler

import (
	"app/internal/sections/models"
	"app/internal/sections/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentHandler(repo *repository.DepartmentRepository) *DepartmentHandler {
	return &DepartmentHandler{
		repo: repo,
	}
}

func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	deparments, err := h.repo.GetAllDepartments()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deparments)
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var input models.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato Inválido"})
		return
	}
	department, err := h.repo.CreateDepartment(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)
}

func (h *DepartmentHandler) EditDepartment(c *gin.Context){
	
	id := c.Param("id")
	var input models.EditDepartmentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	updated, err := h.repo.EditDepartment(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

