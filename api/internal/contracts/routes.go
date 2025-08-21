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
		contractGroup.GET("/", contractHandler.GetAllContracts)
		contractGroup.GET("/:id", contractHandler.GetContractByID)
		contractGroup.POST("/", contractHandler.CreateContract)
		contractGroup.PUT("/:id", contractHandler.EditContract)
		contractGroup.PUT("/values/:id", contractHandler.EditContractValues)
		contractGroup.PUT("/infos/:id", contractHandler.EditContractInfos)
		contractGroup.PUT("/dates/:id", contractHandler.EditContractDates)
		contractGroup.PUT("/discount/:id", contractHandler.EditContractDiscount)
		contractGroup.PUT("/rh/:id", contractHandler.EditContractRhInfo)
		contractGroup.DELETE("/:id", contractHandler.DeactivateContract)
		contractGroup.PATCH("/:id/reactivate", contractHandler.ReactivateContract)
	}
}
