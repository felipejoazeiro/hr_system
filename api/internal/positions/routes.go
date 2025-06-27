package positions


func RegisterRoutes(rg *gin.RouterGroup){
	positionsGroup := rg.Group("/positions")
	{
		positionsGroup.GET("/", handler.PositionsHandler)
		positionsGroup.POST("/", handler.CreatePosition)
		positionsGroup.PUT("/:id", handler.EditPosotion)
		positionsGroup.PUT("/:id/deactivate", handler.DeactivatePosition)
		positionsGroup.PUT("/:id/reactivate", handler.ReactivePosition)
	}
}