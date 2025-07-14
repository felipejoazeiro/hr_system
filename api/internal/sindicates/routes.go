package sindicates

func RegisterRoutes(rg *gin.RouterGroup) {
	repo := repository.NewSindicateRepository(db.DB.DB)
	sindicateHandler := handler.NewSindicateHandler(repo)
	sindicateGroup := rg.Group("/sindicate")
	{
		sindicateGroup.GET("/sindicate")
		sindicateGroup.POST("/sindicate")
		sindicateGroup.PUT("/sindicate/:id")
		sindicateGroup.PATCH("/sindicate/:id/deactivate")
		sindicateGroup.PATCH("/sindicate/:id/reactivate")
	}
}