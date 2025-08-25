package routes

import (
	"app/internal/contracts"
	"app/internal/hr"
	"app/internal/positions"
	"app/internal/sections"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	api:= router.Group("./api")
	{
		hr.RegisterRoutes(api)
		sections.RegisterRoutes(api)
		sindicates.RegisterRoutes(api)
		positions.RegisterRoutes(api)
		contracts.RegisterRoutes(api)
	}
	return router
}