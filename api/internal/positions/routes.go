package positions

import (
	"app/internal/positions/handler"
	"app/internal/positions/repository"
	"app/pkg/db"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	repo := repository.NewPositionRepository(db.DB.DB)
	positionHandler := handler.NewPositionHandler(repo)
	positionsGroup := rg.Group("/positions")
	{
		positionsGroup.GET("/", positionHandler.GetAllPositions)
		positionsGroup.POST("/", positionHandler.CreatePosition)
		positionsGroup.PUT("/:id", positionHandler.EditPosition)
		positionsGroup.PUT("/:id/deactivate", positionHandler.DeactivatePosition)
		positionsGroup.PUT("/:id/reactivate", positionHandler.ReactivePosition)
	}
}
