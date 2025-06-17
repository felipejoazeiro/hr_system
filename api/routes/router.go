package routes

import (
	"github.com/gin-gonic/gin"
	"system/internal/hr"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	api:= router.Group("./api")
	{
		hr.RegisterRoutes(api)
	}
	return router
}