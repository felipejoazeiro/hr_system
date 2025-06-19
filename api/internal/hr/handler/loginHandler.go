package handler

import (
	"app/internal/auth"
	"app/internal/hr/models"
	"app/internal/hr/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context){
	var creds models.EmployeeAccessModel
	if err := c.ShouldBindJSON(&creds); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato Inválido"})
		return
	}

	user, err := repository.Repository.GetUserWithPosition(creds.Login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	if err := verifyPassword(user.Password, creds.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	token, err := auth.GenerateToken(user.Login, user.Position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"position":   user.Position,
		"expires_in": auth.GetTokenExpiration(),
	})
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}