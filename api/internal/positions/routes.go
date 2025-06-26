package positions


func RegisterRoutes(rg *gin.RouterGroup){
	positionsGroup := rg.Group("/employees")
	{
		positionsGroup.GET("/", handler.PositionsHandler)
		positionsGroup.POST("/", handler.CreatePosition)
		positionsGroup.PUT("/:id", handler.EditPosotion)
		positionsGroup.PUT("/:id", handler.DeactivatePosition)
		positionsGroup.PUT("/:id", handler.ReactivePosition)
	}
}