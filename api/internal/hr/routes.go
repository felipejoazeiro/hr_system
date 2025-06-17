package hr

import "github/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup){
	hrGroup := rg.Group("funcionarios")
	{
		hrGroup.POST("", CreateEmployeeHandler)
		// futuramente:
		// hrGroup.GET("", listEmployeeHandler)
		// hrGroup.PUT("/:id", editEmployeeHandler)
		// hrGroup.DELETE("/:id", deleteEmployeeHandler)
	}
}