package hr

import (
	"github.com/gin-gonic/gin"
	"app/internal/hr/handler"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	hrGroup := rg.Group("employees")
	{
		hrGroup.POST("createEmployee", handler.CreateEmployeeHandler)
		hrGroup.GET("/login", handler.LoginHandler)
		hrGroup.GET("/employees", handler.ListEmployeeHandler)
		hrGroup.DELETE("/:id/deactivate", handler.DeactivateEmployeeHandler)
		hrGroup.PUT("/:id", handler.EditEmployeeHandler)
		hrGroup.PATCH("/:id", handler.GetEmployeeByIDHandler) 
		// futuramente:
		// hrGroup.PUT("/:id", editEmployeeHandler)
	}
}
