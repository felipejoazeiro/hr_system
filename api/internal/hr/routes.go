package hr

import (
	"github.com/gin-gonic/gin"
	"app/internal/hr/handler"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	hrGroup := rg.Group("employees")
	{
		hrGroup.POST("", handler.CreateEmployeeHandler)
		hrGroup.GET("", handler.ListEmployeeHandler)
		hrGroup.PATCH("/:id/deactivate", handler.DeactivateEmployeeHandler)
		// futuramente:
		// hrGroup.GET("", listEmployeeHandler)
		// hrGroup.PUT("/:id", editEmployeeHandler)
	}
}
