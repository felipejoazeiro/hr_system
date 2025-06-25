package routes

import (
	"github.com/gin-gonic/gin"
	"app/internal/hr"
	"app/internal/sections"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	api:= router.Group("./api")
	{
		hr.RegisterRoutes(api)
		sections.RegisterRoutes(api)
	}
	return router
}