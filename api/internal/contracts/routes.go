package contracts

import (
	"app/internal/contracts/handler"
	"app/internal/contracts/repository"
	"app/pkg/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	repo := repository.NewContractsRepository(db.DB.DB)
	contractHandler := handler.NewContractHandler(repo)
	contractGroup := rg.Group("/contracts")
	{
		contractGroup.POST("/", contractHandler.CreateContract)
	}
}
