package hr

import (
	"github.com/gin-gonic/gin"
	"app/internal/hr/handler"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	hrGroup := rg.Group("funcionarios")
	{
		hrGroup.POST("", handler.CreateEmployeeHandler)
		hrGroup.GET("/login", handler.LoginHandler)
		// futuramente:
		// hrGroup.GET("", listEmployeeHandler)
		// hrGroup.PUT("/:id", editEmployeeHandler)
		// hrGroup.DELETE("/:id", deleteEmployeeHandler)
	}
}
